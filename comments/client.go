package comments

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/mathisbot/notionary-go/blocks"
	notionhttp "github.com/mathisbot/notionary-go/http"
)

const defaultCommentsPageSize = 100

type Client struct {
	http *notionhttp.Client
}

func New(http *notionhttp.Client) *Client {
	return &Client{http: http}
}

func (c *Client) StreamByBlock(ctx context.Context, blockID string) (<-chan Comment, <-chan error) {
	comments := make(chan Comment)
	errc := make(chan error, 1)

	go func() {
		defer close(comments)
		defer close(errc)

		var cursor string
		for {
			resp, err := c.listCommentsPage(ctx, blockID, cursor, defaultCommentsPageSize)
			if err != nil {
				errc <- err
				return
			}

			for _, comment := range resp.Results {
				select {
				case comments <- comment:
				case <-ctx.Done():
					errc <- ctx.Err()
					return
				}
			}

			if !resp.HasMore || resp.NextCursor == nil || *resp.NextCursor == "" {
				return
			}
			cursor = *resp.NextCursor
		}
	}()

	return comments, errc
}

func (c *Client) CreateForPage(ctx context.Context, pageID string, richText []blocks.RichText) (*Comment, error) {
	request := NewPageCommentRequest(pageID, richText)
	return c.create(ctx, request)
}

func (c *Client) CreateForBlock(ctx context.Context, blockID string, richText []blocks.RichText) (*Comment, error) {
	request := NewBlockCommentRequest(blockID, richText)
	return c.create(ctx, request)
}

func (c *Client) CreateForDiscussion(ctx context.Context, discussionID string, richText []blocks.RichText) (*Comment, error) {
	request := NewDiscussionCommentRequest(discussionID, richText)
	return c.create(ctx, request)
}

func (c *Client) create(ctx context.Context, request CommentCreateRequest) (*Comment, error) {
	var comment Comment
	if err := c.http.Post(ctx, "/comments", request, &comment); err != nil {
		return nil, err
	}
	return &comment, nil
}

func (c *Client) listCommentsPage(ctx context.Context, blockID, startCursor string, pageSize int) (*commentListResponse, error) {
	path := "/comments?block_id=" + url.QueryEscape(blockID)
	if pageSize > 0 {
		path += "&page_size=" + strconv.Itoa(pageSize)
	}
	if startCursor != "" {
		path += "&start_cursor=" + url.QueryEscape(startCursor)
	}

	var resp commentListResponse
	if err := c.http.Get(ctx, path, &resp); err != nil {
		return nil, fmt.Errorf("list comments for block %q: %w", blockID, err)
	}

	return &resp, nil
}
