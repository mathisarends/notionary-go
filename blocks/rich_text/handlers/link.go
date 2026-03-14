package handlers

import (
	"regexp"

	"github.com/mathisbot/notionary-go/blocks"
)

var LinkPattern = regexp.MustCompile(`\[(.+?)\]\((https?://[^\s)]+)\)`)

type LinkHandler struct{}

func (LinkHandler) Tag() string { return "[" }

func LinkToMarkdown(text, url, prefix, middle, suffix string) string {
	if text == "" {
		return ""
	}
	if url == "" {
		return text
	}
	if prefix == "" {
		prefix = "["
	}
	if middle == "" {
		middle = "]("
	}
	if suffix == "" {
		suffix = ")"
	}
	return prefix + text + middle + url + suffix
}

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
