package file_upload

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	notionhttp "github.com/mathisbot/notionary-go/http"
)

type Client struct {
	http   *notionhttp.Client
	config UploadConfig
}

func New(http *notionhttp.Client) *Client {
	return NewWithConfig(http, DefaultUploadConfig())
}

func NewWithConfig(http *notionhttp.Client, cfg UploadConfig) *Client {
	if cfg.SinglePartMaxSize <= 0 {
		cfg.SinglePartMaxSize = NotionSinglePartMaxSize
	}
	if cfg.MultipartChunkSize <= 0 {
		cfg.MultipartChunkSize = DefaultMultipartChunkSize
	}
	if cfg.FilenameByteLimit <= 0 {
		cfg.FilenameByteLimit = DefaultFilenameByteLimit
	}
	if cfg.PollInterval <= 0 {
		cfg.PollInterval = DefaultPollInterval
	}
	if cfg.MaxUploadWait <= 0 {
		cfg.MaxUploadWait = DefaultMaxUploadWait
	}

	return &Client{http: http, config: cfg}
}

func (c *Client) UploadFile(ctx context.Context, filePath string, opts *UploadOptions) (*FileUpload, error) {
	resolvedPath, err := c.resolvePath(filePath)
	if err != nil {
		return nil, err
	}

	stat, err := os.Stat(resolvedPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("file not found: %s", resolvedPath)
		}
		return nil, fmt.Errorf("stat file: %w", err)
	}

	if stat.IsDir() {
		return nil, fmt.Errorf("path is a directory, expected file: %s", resolvedPath)
	}

	options := c.normalizeOptions(opts, filepath.Base(resolvedPath))
	if err := c.validateFilename(options.Filename); err != nil {
		return nil, err
	}

	contentType := options.ContentType
	if contentType == "" {
		contentType = guessContentType(options.Filename)
	}

	if stat.Size() <= c.config.SinglePartMaxSize {
		content, err := os.ReadFile(resolvedPath)
		if err != nil {
			return nil, fmt.Errorf("read file: %w", err)
		}
		return c.uploadSinglePart(ctx, content, options.Filename, contentType, options)
	}

	return c.uploadMultiPartFromFile(ctx, resolvedPath, stat.Size(), options.Filename, contentType, options)
}

func (c *Client) UploadBytes(ctx context.Context, content []byte, filename string, opts *UploadOptions) (*FileUpload, error) {
	if len(content) == 0 {
		return nil, errors.New("file content is empty")
	}

	options := c.normalizeOptions(opts, filename)
	if err := c.validateFilename(options.Filename); err != nil {
		return nil, err
	}

	contentType := options.ContentType
	if contentType == "" {
		contentType = guessContentType(options.Filename)
	}

	if int64(len(content)) <= c.config.SinglePartMaxSize {
		return c.uploadSinglePart(ctx, content, options.Filename, contentType, options)
	}

	return c.uploadMultiPartFromBytes(ctx, content, options.Filename, contentType, options)
}

func (c *Client) Get(ctx context.Context, uploadID string) (*FileUpload, error) {
	var upload FileUpload
	if err := c.http.Get(ctx, "/file_uploads/"+uploadID, &upload); err != nil {
		return nil, err
	}
	return &upload, nil
}

func (c *Client) GetStatus(ctx context.Context, uploadID string) (FileUploadStatus, error) {
	upload, err := c.Get(ctx, uploadID)
	if err != nil {
		return "", err
	}
	return upload.Status, nil
}

func (c *Client) WaitForCompletion(ctx context.Context, uploadID string, timeout time.Duration) (*FileUpload, error) {
	if timeout <= 0 {
		timeout = c.config.MaxUploadWait
	}

	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	ticker := time.NewTicker(c.config.PollInterval)
	defer ticker.Stop()

	for {
		upload, err := c.Get(ctxWithTimeout, uploadID)
		if err != nil {
			if errors.Is(ctxWithTimeout.Err(), context.DeadlineExceeded) {
				return nil, &UploadTimeoutError{UploadID: uploadID, Timeout: timeout}
			}
			return nil, &UploadFailedError{UploadID: uploadID, Reason: err.Error()}
		}

		switch upload.Status {
		case FileUploadStatusUploaded:
			return upload, nil
		case FileUploadStatusFailed, FileUploadStatusExpired:
			return nil, &UploadFailedError{UploadID: uploadID, Reason: "upload status: " + string(upload.Status)}
		}

		select {
		case <-ctxWithTimeout.Done():
			return nil, &UploadTimeoutError{UploadID: uploadID, Timeout: timeout}
		case <-ticker.C:
		}
	}
}

func (c *Client) List(ctx context.Context, query ListUploadsQuery) (*ListUploadsResponse, error) {
	values := url.Values{}
	if query.PageSize > 0 {
		if query.PageSize > 100 {
			query.PageSize = 100
		}
		values.Set("page_size", strconv.Itoa(query.PageSize))
	}
	if query.StartCursor != "" {
		values.Set("start_cursor", query.StartCursor)
	}
	if query.Status != "" {
		values.Set("status", string(query.Status))
	}
	if query.Archived != nil {
		values.Set("archived", strconv.FormatBool(*query.Archived))
	}

	path := "/file_uploads"
	if encoded := values.Encode(); encoded != "" {
		path += "?" + encoded
	}

	var resp ListUploadsResponse
	if err := c.http.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) ListAll(ctx context.Context, query ListUploadsQuery) ([]FileUpload, error) {
	var all []FileUpload
	for {
		resp, err := c.List(ctx, query)
		if err != nil {
			return nil, err
		}
		all = append(all, resp.Results...)
		if !resp.HasMore {
			return all, nil
		}
		query.StartCursor = resp.NextCursor
	}
}

func (c *Client) uploadSinglePart(
	ctx context.Context,
	content []byte,
	filename string,
	contentType string,
	options UploadOptions,
) (*FileUpload, error) {
	upload, err := c.createUpload(ctx, filename, contentType, UploadModeSinglePart, nil, int64(len(content)))
	if err != nil {
		return nil, err
	}

	if err := c.sendPart(ctx, upload.ID, filename, content, 0); err != nil {
		return nil, err
	}

	if !options.WaitForCompletion {
		return upload, nil
	}

	return c.WaitForCompletion(ctx, upload.ID, options.Timeout)
}

func (c *Client) uploadMultiPartFromBytes(
	ctx context.Context,
	content []byte,
	filename string,
	contentType string,
	options UploadOptions,
) (*FileUpload, error) {
	partCount := c.calculatePartCount(int64(len(content)))
	upload, err := c.createUpload(ctx, filename, contentType, UploadModeMultiPart, &partCount, int64(len(content)))
	if err != nil {
		return nil, err
	}

	for i, part := 0, 1; i < len(content); i, part = i+c.config.MultipartChunkSize, part+1 {
		end := i + c.config.MultipartChunkSize
		if end > len(content) {
			end = len(content)
		}
		if err := c.sendPart(ctx, upload.ID, filename, content[i:end], part); err != nil {
			return nil, err
		}
	}

	if err := c.completeUpload(ctx, upload.ID); err != nil {
		return nil, err
	}

	if !options.WaitForCompletion {
		return upload, nil
	}

	return c.WaitForCompletion(ctx, upload.ID, options.Timeout)
}

func (c *Client) uploadMultiPartFromFile(
	ctx context.Context,
	resolvedPath string,
	fileSize int64,
	filename string,
	contentType string,
	options UploadOptions,
) (*FileUpload, error) {
	partCount := c.calculatePartCount(fileSize)
	upload, err := c.createUpload(ctx, filename, contentType, UploadModeMultiPart, &partCount, fileSize)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(resolvedPath)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	buf := make([]byte, c.config.MultipartChunkSize)
	part := 1
	for {
		n, readErr := io.ReadFull(file, buf)
		if errors.Is(readErr, io.EOF) {
			break
		}
		if readErr != nil && !errors.Is(readErr, io.ErrUnexpectedEOF) {
			return nil, fmt.Errorf("read chunk %d/%d: %w", part, partCount, readErr)
		}
		if n > 0 {
			if err := c.sendPart(ctx, upload.ID, filename, buf[:n], part); err != nil {
				return nil, err
			}
			part++
		}
		if errors.Is(readErr, io.ErrUnexpectedEOF) {
			break
		}
	}

	if err := c.completeUpload(ctx, upload.ID); err != nil {
		return nil, err
	}

	if !options.WaitForCompletion {
		return upload, nil
	}

	return c.WaitForCompletion(ctx, upload.ID, options.Timeout)
}

func (c *Client) createUpload(
	ctx context.Context,
	filename string,
	contentType string,
	mode UploadMode,
	numberOfParts *int,
	contentLength int64,
) (*FileUpload, error) {
	req := createUploadRequest{
		Filename:      filename,
		ContentType:   contentType,
		ContentLength: &contentLength,
		Mode:          mode,
		NumberOfParts: numberOfParts,
	}

	var upload FileUpload
	if err := c.http.Post(ctx, "/file_uploads", req, &upload); err != nil {
		return nil, err
	}

	return &upload, nil
}

func (c *Client) sendPart(
	ctx context.Context,
	uploadID string,
	filename string,
	content []byte,
	partNumber int,
) error {
	fields := map[string]string{}
	if partNumber > 0 {
		fields["part_number"] = strconv.Itoa(partNumber)
	}

	var upload FileUpload
	err := c.http.PostMultipart(
		ctx,
		"/file_uploads/"+uploadID+"/send",
		"file",
		filename,
		content,
		fields,
		&upload,
	)
	if err != nil {
		return &UploadFailedError{UploadID: uploadID, Reason: err.Error()}
	}

	return nil
}

func (c *Client) completeUpload(ctx context.Context, uploadID string) error {
	var upload FileUpload
	if err := c.http.Post(ctx, "/file_uploads/"+uploadID+"/complete", struct{}{}, &upload); err != nil {
		return &UploadFailedError{UploadID: uploadID, Reason: err.Error()}
	}
	return nil
}

func (c *Client) normalizeOptions(opts *UploadOptions, fallbackName string) UploadOptions {
	resolved := UploadOptions{
		Filename:          fallbackName,
		WaitForCompletion: true,
	}
	if opts == nil {
		return resolved
	}
	if opts.Filename != "" {
		resolved.Filename = opts.Filename
	}
	resolved.ContentType = opts.ContentType
	resolved.Timeout = opts.Timeout
	resolved.WaitForCompletion = opts.WaitForCompletion
	return resolved
}

func (c *Client) resolvePath(path string) (string, error) {
	trimmed := strings.TrimSpace(path)
	if trimmed == "" {
		return "", errors.New("file path is required")
	}

	resolved := trimmed
	if !filepath.IsAbs(resolved) && c.config.BaseUploadPath != "" {
		resolved = filepath.Join(c.config.BaseUploadPath, resolved)
	}

	abs, err := filepath.Abs(resolved)
	if err != nil {
		return "", fmt.Errorf("resolve file path: %w", err)
	}
	return abs, nil
}

func (c *Client) calculatePartCount(fileSize int64) int {
	chunk := int64(c.config.MultipartChunkSize)
	return int((fileSize + chunk - 1) / chunk)
}

func (c *Client) validateFilename(filename string) error {
	if strings.TrimSpace(filename) == "" {
		return errors.New("filename is required")
	}
	if len([]byte(filename)) > c.config.FilenameByteLimit {
		return fmt.Errorf("filename too long: %d bytes (max %d)", len([]byte(filename)), c.config.FilenameByteLimit)
	}
	return nil
}

func guessContentType(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	if ext == "" {
		return ""
	}
	return mime.TypeByExtension(ext)
}

type UploadFailedError struct {
	UploadID string
	Reason   string
}

func (e *UploadFailedError) Error() string {
	if e.Reason == "" {
		return fmt.Sprintf("file upload %s failed", e.UploadID)
	}
	return fmt.Sprintf("file upload %s failed: %s", e.UploadID, e.Reason)
}

type UploadTimeoutError struct {
	UploadID string
	Timeout  time.Duration
}

func (e *UploadTimeoutError) Error() string {
	return fmt.Sprintf("file upload %s did not complete in %s", e.UploadID, e.Timeout)
}
