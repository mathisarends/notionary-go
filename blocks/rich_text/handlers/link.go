package handlers

import (
	"regexp"

	"github.com/mathisbot/notionary-go/blocks"
)

var LinkPattern = regexp.MustCompile(`\[(.+?)\]\((https?://[^\s)]+)\)`)

type LinkHandler struct{}

func (LinkHandler) Tag() string { return "[" }

func (LinkHandler) Handle(match []string) blocks.RichText {
	url := match[2]
	return blocks.RichText{
		Type:      blocks.RichTextTypeText,
		PlainText: match[1],
		Text: &blocks.TextContent{
			Content: match[1],
			Link: &struct {
				URL string `json:"url"`
			}{URL: url},
		},
	}
}
