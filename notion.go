package notionary

import (
	"github.com/mathisbot/notionary-go/database"
	"github.com/mathisbot/notionary-go/datasource"
	notionhttp "github.com/mathisbot/notionary-go/http"
	"github.com/mathisbot/notionary-go/page"
)

type Client struct {
	Pages       *page.Client
	Databases   *database.Client
	DataSources *datasource.Client
}

func New(token string) *Client {
	http := notionhttp.New(token)
	return &Client{
		Pages:       page.New(http),
		Databases:   database.New(http),
		DataSources: datasource.New(http),
	}
}