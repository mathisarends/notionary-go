package markdown

import (
	"fmt"
	"strings"

	"github.com/mathisbot/notionary-go/blocks"
)

func Render(bs []blocks.Block) string {
	var sb strings.Builder
	for _, b := range bs {
		line := renderBlock(b)
		if line != "" {
			sb.WriteString(line)
			sb.WriteString("\n")
		}
	}
	return strings.TrimSpace(sb.String())
}

func renderBlock(b blocks.Block) string {
	switch v := b.(type) {
	case *blocks.Heading1Block:
		return "# " + plainText(v.Heading1.RichText)
	case *blocks.Heading2Block:
		return "## " + plainText(v.Heading2.RichText)
	case *blocks.Heading3Block:
		return "### " + plainText(v.Heading3.RichText)
	case *blocks.ParagraphBlock:
		return plainText(v.Paragraph.RichText)
	case *blocks.BulletedListItemBlock:
		return "- " + plainText(v.BulletedListItem.RichText)
	case *blocks.NumberedListItemBlock:
		return "1. " + plainText(v.NumberedListItem.RichText)
	case *blocks.ToDoBlock:
		if v.ToDo.Checked {
			return "- [x] " + plainText(v.ToDo.RichText)
		}
		return "- [ ] " + plainText(v.ToDo.RichText)
	case *blocks.DividerBlock:
		return "---"
	case *blocks.CodeBlock:
		return fmt.Sprintf("```%s\n%s\n```", v.Code.Language, plainText(v.Code.RichText))
	case *blocks.QuoteBlock:
		return "> " + plainText(v.Quote.RichText)
	default:
		return ""
	}
}

func plainText(rts []blocks.RichText) string {
	var sb strings.Builder
	for _, rt := range rts {
		sb.WriteString(rt.PlainText)
	}
	return sb.String()
}