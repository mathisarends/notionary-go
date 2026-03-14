package page

import (
	"context"
	"fmt"

	notionhttp "github.com/mathisbot/notionary-go/http"
	search "github.com/mathisbot/notionary-go/search"
)

type Client struct {
	http *notionhttp.Client
}

func New(http *notionhttp.Client) *Client {
	return &Client{http: http}
}

func (c *Client) Get(ctx context.Context, id string) (*Page, error) {
	var page Page
	if err := c.http.Get(ctx, "/pages/"+id, &page); err != nil {
		return nil, err
	}
	return &page, nil
}

func (c *Client) FindByTitle(ctx context.Context, title string) (*Page, error) {
	var cursor string

	for {
		var resp search.SearchResponse
		err := c.http.Post(ctx, "/search", search.SearchRequest{
			Query:       title,
			Filter:      &search.SearchFilter{Value: "page", Property: "object"},
			StartCursor: cursor,
		}, &resp)
		if err != nil {
			return nil, err
		}

		for _, result := range resp.Results {
			if result.PlainTitle() == title {
				return c.Get(ctx, result.ID)
			}
		}

		if !resp.HasMore {
			break
		}
		cursor = resp.NextCursor
	}

	return nil, fmt.Errorf("page %q not found", title)
}

func (c *Client) Stream(ctx context.Context, config search.WorkspaceQueryConfig) (<-chan *Page, <-chan error) {
	pages := make(chan *Page)
	errc := make(chan error, 1)

	config = config.WithPagesOnly()

	go func() {
		defer close(pages)
		defer close(errc)

		var cursor string
		for {
			var resp search.SearchResponse
			err := c.http.Post(ctx, "/search", config.WithStartCursor(cursor).ToAPIParams(), &resp)
			if err != nil {
				errc <- err
				return
			}

			for _, result := range resp.Results {
				page, err := c.Get(ctx, result.ID)
				if err != nil {
					errc <- err
					return
				}
				select {
				case pages <- page:
				case <-ctx.Done():
					errc <- ctx.Err()
					return
				}
			}

			if !resp.HasMore {
				return
			}
			cursor = resp.NextCursor
		}
	}()

	return pages, errc
}
