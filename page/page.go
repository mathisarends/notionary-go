package page

import (
	"context"
	"fmt"

	"github.com/mathisbot/notionary-go/blocks"
	"github.com/mathisbot/notionary-go/blocks/markdown"
	"github.com/mathisbot/notionary-go/comments"
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

type Page struct {
	ID             string         `json:"id"`
	Object         string         `json:"object"`
	CreatedTime    string         `json:"created_time"`
	LastEditedTime string         `json:"last_edited_time"`
	Archived       bool           `json:"archived"`
	Properties     map[string]any `json:"properties"`

	http     *notionhttp.Client
	comments *comments.Client
	pending  []op
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