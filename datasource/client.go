package datasource

import notionhttp "github.com/mathisbot/notionary-go/http"

type Client struct {
	http *notionhttp.Client
}

func New(http *notionhttp.Client) *Client {
	return &Client{http: http}
}