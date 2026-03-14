package notionary

import (
	"github.com/mathisbot/notionary-go/comments"
	"github.com/mathisbot/notionary-go/database"
	"github.com/mathisbot/notionary-go/datasource"
	"github.com/mathisbot/notionary-go/file_upload"
	notionhttp "github.com/mathisbot/notionary-go/http"
	"github.com/mathisbot/notionary-go/page"
	"github.com/mathisbot/notionary-go/user"
)

type Client struct {
	Pages       *page.Client
	Comments    *comments.Client
	Databases   *database.Client
	DataSources *datasource.Client
	FileUploads *file_upload.Client
	Users       *user.Client
}

func New(token string) *Client {
	http := notionhttp.New(token)
	return &Client{
		Pages:       page.New(http),
		Comments:    comments.New(http),
		Databases:   database.New(http),
		DataSources: datasource.New(http),
		FileUploads: file_upload.New(http),
		Users:       user.New(http),
	}
}
