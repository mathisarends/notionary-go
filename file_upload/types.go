package file_upload

import "time"

type UploadMode string

const (
	UploadModeSinglePart UploadMode = "single_part"
	UploadModeMultiPart  UploadMode = "multi_part"
)

type FileUploadStatus string

const (
	FileUploadStatusPending  FileUploadStatus = "pending"
	FileUploadStatusUploaded FileUploadStatus = "uploaded"
	FileUploadStatusFailed   FileUploadStatus = "failed"
	FileUploadStatusExpired  FileUploadStatus = "expired"
)

const (
	NotionSinglePartMaxSize   int64 = 20 * 1024 * 1024
	DefaultMultipartChunkSize       = int(NotionSinglePartMaxSize)
	DefaultFilenameByteLimit        = 900
	DefaultPollInterval             = 2 * time.Second
	DefaultMaxUploadWait            = 2 * time.Minute
)

type UploadConfig struct {
	SinglePartMaxSize  int64
	MultipartChunkSize int
	FilenameByteLimit  int
	PollInterval       time.Duration
	MaxUploadWait      time.Duration
	BaseUploadPath     string
}

func DefaultUploadConfig() UploadConfig {
	return UploadConfig{
		SinglePartMaxSize:  NotionSinglePartMaxSize,
		MultipartChunkSize: DefaultMultipartChunkSize,
		FilenameByteLimit:  DefaultFilenameByteLimit,
		PollInterval:       DefaultPollInterval,
		MaxUploadWait:      DefaultMaxUploadWait,
	}
}

type UploadOptions struct {
	Filename          string
	ContentType       string
	WaitForCompletion bool
	Timeout           time.Duration
}

type ListUploadsQuery struct {
	PageSize    int
	StartCursor string
	Status      FileUploadStatus
	Archived    *bool
}

type FileUpload struct {
	ID             string           `json:"id"`
	CreatedTime    string           `json:"created_time"`
	LastEditedTime string           `json:"last_edited_time"`
	ExpiryTime     *string          `json:"expiry_time,omitempty"`
	UploadURL      *string          `json:"upload_url,omitempty"`
	Archived       bool             `json:"archived"`
	Status         FileUploadStatus `json:"status"`
	Filename       *string          `json:"filename,omitempty"`
	ContentType    *string          `json:"content_type,omitempty"`
	ContentLength  *int64           `json:"content_length,omitempty"`
	RequestID      *string          `json:"request_id,omitempty"`
}

type ListUploadsResponse struct {
	Results    []FileUpload `json:"results"`
	NextCursor string       `json:"next_cursor"`
	HasMore    bool         `json:"has_more"`
}

type createUploadRequest struct {
	Filename      string     `json:"filename"`
	ContentType   string     `json:"content_type,omitempty"`
	ContentLength *int64     `json:"content_length,omitempty"`
	Mode          UploadMode `json:"mode"`
	NumberOfParts *int       `json:"number_of_parts,omitempty"`
}
