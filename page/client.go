package page

import (
	"context"
	"fmt"

	notionhttp "github.com/mathisbot/notionary-go/http"
	"github.com/mathisbot/notionary-go/shared"
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
	page.http = c.http
	return &page, nil
}

func (c *Client) FindByTitle(ctx context.Context, title string) (*Page, error) {
	var cursor string

	for {
		var resp shared.SearchResponse
		err := c.http.Post(ctx, "/search", shared.SearchRequest{
			Query:       title,
			Filter:      &shared.SearchFilter{Value: "page", Property: "object"},
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