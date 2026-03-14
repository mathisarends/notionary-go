package markdown

import (
	"github.com/mathisbot/notionary-go/blocks"
	richtext "github.com/mathisbot/notionary-go/blocks/rich_text"
)

var defaultRichTextConverter = richtext.NewDefaultConverter()

func toRichText(text string) []blocks.RichText {
	return defaultRichTextConverter.ToRichText(text)
}
