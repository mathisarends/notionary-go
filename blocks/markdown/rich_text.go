package markdown

import "github.com/mathisbot/notionary-go/blocks"

func toRichText(text string) []blocks.RichText {
	return []blocks.RichText{
		{
			Type: blocks.RichTextTypeText,
			Text: &blocks.TextContent{Content: text},
			PlainText: text,
		},
	}
}