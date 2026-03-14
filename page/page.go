package page

import (
	"context"
	"fmt"

	"github.com/mathisbot/notionary-go/blocks"
	"github.com/mathisbot/notionary-go/blocks/markdown"
	notionhttp "github.com/mathisbot/notionary-go/http"
)

type Page struct {
	ID             string         `json:"id"`
	Object         string         `json:"object"`
	CreatedTime    string         `json:"created_time"`
	LastEditedTime string         `json:"last_edited_time"`
	Archived       bool           `json:"archived"`
	Properties     map[string]any `json:"properties"`

	http    *notionhttp.Client
	pending []any
}

func (p *Page) Content(ctx context.Context) (string, error) {
	var allBlocks []blocks.Block
	var cursor string

	for {
		var resp blocks.BlockList
		path := fmt.Sprintf("/blocks/%s/children", p.ID)
		if cursor != "" {
			path += "?start_cursor=" + cursor
		}
		if err := p.http.Get(ctx, path, &resp); err != nil {
			return "", err
		}
		allBlocks = append(allBlocks, resp.Results...)
		if !resp.HasMore {
			break
		}
		cursor = *resp.NextCursor
	}

	return markdown.Render(allBlocks), nil
}

func (p *Page) Append(md string) {
	_ = p.AppendMarkdown(md)
}

func (p *Page) AppendMarkdown(md string) error {
	parsed, err := markdown.Parse(md)
	if err != nil {
		return err
	}
	p.pending = append(p.pending, parsed...)
	return nil
}

func (p *Page) Commit(ctx context.Context) error {
	if len(p.pending) == 0 {
		return nil
	}
	payload := map[string]any{"children": p.pending}
	if err := p.http.Patch(ctx, fmt.Sprintf("/blocks/%s/children", p.ID), payload, nil); err != nil {
		return err
	}
	p.pending = nil
	return nil
}
