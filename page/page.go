package page

import (
	"context"
	"fmt"

	"github.com/mathisbot/notionary-go/blocks"
	"github.com/mathisbot/notionary-go/blocks/markdown"
	"github.com/mathisbot/notionary-go/comments"
	"github.com/mathisbot/notionary-go/entity"
	notionhttp "github.com/mathisbot/notionary-go/http"
)

type op interface {
	apply(ctx context.Context, p *Page) error
}

type appendBlocksOp struct {
	children []any
}

type clearContentOp struct{}

type setTitleOp struct {
	title string
}

type postCommentOp struct {
	richText []blocks.RichText
}

type postReplyOp struct {
	discussionID string
	richText     []blocks.RichText
}

type pageResponseDTO struct {
	entity.EntityResponseDto
	Object     string         `json:"object"`
	Archived   bool           `json:"archived"`
	Properties map[string]any `json:"properties"`
}

type pageMetadataClient struct {
	http   *notionhttp.Client
	pageID string
}

type Page struct {
	entity.Entity `json:"-"`
	Object        string         `json:"object"`
	Archived      bool           `json:"archived"`
	Properties    map[string]any `json:"properties"`

	http     *notionhttp.Client
	comments *comments.Client
	pending  []op
}

func newPageFromDTO(dto pageResponseDTO, http *notionhttp.Client) *Page {
	meta := pageMetadataClient{http: http, pageID: dto.ID}
	return &Page{
		Entity:     entity.New(dto.EntityResponseDto, meta, http),
		Object:     dto.Object,
		Archived:   dto.Archived,
		Properties: dto.Properties,
		http:       http,
		comments:   comments.New(http),
	}
}

func (p *Page) AppendMarkdown(md string) error {
	converter := markdown.NewConverter()
	parsed, err := converter.ToBlocks(md)
	if err != nil {
		return err
	}
	p.pending = append(p.pending, appendBlocksOp{children: parsed})
	return nil
}

func (p *Page) ClearContent() {
	p.pending = append(p.pending, clearContentOp{})
}

func (p *Page) ReplaceContent(md string) error {
	p.ClearContent()
	return p.AppendMarkdown(md)
}

func (p *Page) SetTitle(title string) {
	p.pending = append(p.pending, setTitleOp{title: title})
}

func (p *Page) PostTopLevelComment(richText []blocks.RichText) {
	p.pending = append(p.pending, postCommentOp{richText: richText})
}

func (p *Page) PostReplyToDiscussion(discussionID string, richText []blocks.RichText) {
	p.pending = append(p.pending, postReplyOp{discussionID: discussionID, richText: richText})
}

func (p *Page) Commit(ctx context.Context) error {
	for _, o := range p.pending {
		if err := o.apply(ctx, p); err != nil {
			return err
		}
	}
	p.pending = nil
	return nil
}

func (p *Page) Content(ctx context.Context) (string, error) {
	allBlocks, err := p.fetchAllBlocks(ctx)
	if err != nil {
		return "", err
	}
	converter := markdown.NewConverter()
	return converter.ToMarkdown(allBlocks), nil
}

func (p *Page) GetComments(ctx context.Context) ([]comments.Comment, error) {
	var all []comments.Comment
	ch, errc := p.comments.StreamByBlock(ctx, p.ID)
	for c := range ch {
		all = append(all, c)
	}
	if err := <-errc; err != nil {
		return nil, err
	}
	return all, nil
}

func (o appendBlocksOp) apply(ctx context.Context, p *Page) error {
	payload := map[string]any{"children": o.children}
	return p.http.Patch(ctx, fmt.Sprintf("/blocks/%s/children", p.ID), payload, nil)
}

func (o clearContentOp) apply(ctx context.Context, p *Page) error {
	allBlocks, err := p.fetchAllBlocks(ctx)
	if err != nil {
		return err
	}
	for _, b := range allBlocks {
		if err := p.http.Delete(ctx, "/blocks/"+b.GetID(), nil); err != nil {
			return err
		}
	}
	return nil
}

func (o setTitleOp) apply(ctx context.Context, p *Page) error {
	payload := map[string]any{
		"properties": map[string]any{
			"title": map[string]any{
				"title": []map[string]any{
					{"type": "text", "text": map[string]any{"content": o.title}},
				},
			},
		},
	}
	return p.http.Patch(ctx, "/pages/"+p.ID, payload, nil)
}

func (o postCommentOp) apply(ctx context.Context, p *Page) error {
	_, err := p.comments.CreateForPage(ctx, p.ID, o.richText)
	return err
}

func (o postReplyOp) apply(ctx context.Context, p *Page) error {
	_, err := p.comments.CreateForDiscussion(ctx, o.discussionID, o.richText)
	return err
}

func (p *Page) fetchAllBlocks(ctx context.Context) ([]blocks.Block, error) {
	var all []blocks.Block
	var cursor string
	for {
		var resp blocks.BlockList
		path := fmt.Sprintf("/blocks/%s/children", p.ID)
		if cursor != "" {
			path += "?start_cursor=" + cursor
		}
		if err := p.http.Get(ctx, path, &resp); err != nil {
			return nil, err
		}
		all = append(all, resp.Results...)
		if !resp.HasMore {
			break
		}
		cursor = *resp.NextCursor
	}
	return all, nil
}

func (m pageMetadataClient) PatchEmojiIcon(ctx context.Context, emoji string) (*entity.EntityResponseDto, error) {
	return m.patch(ctx, map[string]any{
		"icon": map[string]any{
			"type":  "emoji",
			"emoji": emoji,
		},
	})
}

func (m pageMetadataClient) PatchExternalIcon(ctx context.Context, url string) (*entity.EntityResponseDto, error) {
	return m.patch(ctx, map[string]any{
		"icon": map[string]any{
			"type": "external",
			"external": map[string]any{
				"url": url,
			},
		},
	})
}

func (m pageMetadataClient) PatchIconFromFileUpload(ctx context.Context, fileUploadID string) (*entity.EntityResponseDto, error) {
	return m.patch(ctx, map[string]any{
		"icon": map[string]any{
			"type": "file_upload",
			"file_upload": map[string]any{
				"id": fileUploadID,
			},
		},
	})
}

func (m pageMetadataClient) RemoveIcon(ctx context.Context) error {
	return m.patchEmpty(ctx, map[string]any{"icon": nil})
}

func (m pageMetadataClient) PatchExternalCover(ctx context.Context, url string) (*entity.EntityResponseDto, error) {
	return m.patch(ctx, map[string]any{
		"cover": map[string]any{
			"type": "external",
			"external": map[string]any{
				"url": url,
			},
		},
	})
}

func (m pageMetadataClient) PatchCoverFromFileUpload(ctx context.Context, fileUploadID string) (*entity.EntityResponseDto, error) {
	return m.patch(ctx, map[string]any{
		"cover": map[string]any{
			"type": "file_upload",
			"file_upload": map[string]any{
				"id": fileUploadID,
			},
		},
	})
}

func (m pageMetadataClient) RemoveCover(ctx context.Context) error {
	return m.patchEmpty(ctx, map[string]any{"cover": nil})
}

func (m pageMetadataClient) MoveToTrash(ctx context.Context) (*entity.EntityResponseDto, error) {
	return m.patch(ctx, map[string]any{"in_trash": true})
}

func (m pageMetadataClient) RestoreFromTrash(ctx context.Context) (*entity.EntityResponseDto, error) {
	return m.patch(ctx, map[string]any{"in_trash": false})
}

func (m pageMetadataClient) patch(ctx context.Context, payload map[string]any) (*entity.EntityResponseDto, error) {
	var resp pageResponseDTO
	if err := m.http.Patch(ctx, "/pages/"+m.pageID, payload, &resp); err != nil {
		return nil, err
	}
	return &resp.EntityResponseDto, nil
}

func (m pageMetadataClient) patchEmpty(ctx context.Context, payload map[string]any) error {
	return m.http.Patch(ctx, "/pages/"+m.pageID, payload, nil)
}