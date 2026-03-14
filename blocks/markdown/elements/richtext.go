package elements

import (
	"github.com/mathisbot/notionary-go/blocks"
	richtext "github.com/mathisbot/notionary-go/blocks/rich_text"
)

var richTextCodec = richtext.NewDefaultConverter()

func toRichText(text string) []blocks.RichText {
	return richTextCodec.ToRichText(text)
}

func toMarkdown(richTexts []blocks.RichText) string {
	return richTextCodec.ToMarkdown(richTexts)
}