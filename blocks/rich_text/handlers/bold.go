package handlers

import (
	"regexp"

	"github.com/mathisbot/notionary-go/blocks"
)

var BoldPattern = regexp.MustCompile(`\*\*(.+?)\*\*`)

type BoldHandler struct{}

func (BoldHandler) Tag() string { return "**" }

func (BoldHandler) Handle(match []string) blocks.RichText {
	return blocks.RichText{
		Type:        blocks.RichTextTypeText,
		PlainText:   match[1],
		Text:        &blocks.TextContent{Content: match[1]},
		Annotations: &blocks.Annotations{Bold: true},
	}
}
