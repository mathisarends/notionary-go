package user

import (
	"context"
	"fmt"

	notionhttp "github.com/mathisbot/notionary-go/http"
)

type Client struct {
	http *notionhttp.Client
}

func New(http *notionhttp.Client) *Client {
	return &Client{http: http}
}

func (c *Client) Me(ctx context.Context) (Bot, error) {
	var dto BotUserResponseDto
	if err := c.http.Get(ctx, "/users/me", &dto); err != nil {
		return Bot{}, err
	}
	return NewBot(dto), nil
}

func (c *Client) List(ctx context.Context) ([]UserResponseDto, error) {
	var results []UserResponseDto
	var cursor string

	for {
		var resp NotionUsersListResponse
		url := "/users"
		if cursor != "" {
			url += "?start_cursor=" + cursor
		}
		if err := c.http.Get(ctx, url, &resp); err != nil {
			return nil, err
		}
		results = append(results, resp.Results...)
		if !resp.HasMore {
			break
		}
		cursor = *resp.NextCursor
	}

	return results, nil
}

func (c *Client) FindByID(ctx context.Context, id string) (UserResponseDto, error) {
	var dto UserResponseDto
	if err := c.http.Get(ctx, "/users/"+id, &dto); err != nil {
		return UserResponseDto{}, fmt.Errorf("user %q not found: %w", id, err)
	}
	return dto, nil
}