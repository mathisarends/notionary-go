package markdown

import (
	"fmt"
	"strings"

	"github.com/mathisbot/notionary-go/blocks"
)

type RenderFunc func(blocks.Block) (string, bool)

type BlockRenderer struct {
	handlers []RenderFunc
}

func (r *BlockRenderer) Render(block blocks.Block) (string, bool) {
	for _, h := range r.handlers {
		if s, ok := h(block); ok {
			return s, true
		}
	}
	return "", false
}

func NewRenderer() *BlockRenderer {
	return &BlockRenderer{handlers: []RenderFunc{
		renderHeading1,
		renderHeading2,
		renderHeading3,
		renderParagraph,
		renderToDo,
		renderDivider,
		renderCode,
		renderQuote,
	}}
}

func renderHeading1(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.Heading1Block)
	if !ok {
		return "", false
	}
	return "# " + plainText(b.Heading1.RichText), true
}

func renderHeading2(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.Heading2Block)
	if !ok {
		return "", false
	}
	return "## " + plainText(b.Heading2.RichText), true
}

func renderHeading3(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.Heading3Block)
	if !ok {
		return "", false
	}
	return "### " + plainText(b.Heading3.RichText), true
}

func renderParagraph(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.ParagraphBlock)
	if !ok {
		return "", false
	}
	return plainText(b.Paragraph.RichText), true
}

func renderToDo(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.ToDoBlock)
	if !ok {
		return "", false
	}
	if b.ToDo.Checked {
		return "- [x] " + plainText(b.ToDo.RichText), true
	}
	return "- [ ] " + plainText(b.ToDo.RichText), true
}

func renderDivider(block blocks.Block) (string, bool) {
	if _, ok := block.(*blocks.DividerBlock); !ok {
		return "", false
	}
	return "---", true
}

func renderCode(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.CodeBlock)
	if !ok {
		return "", false
	}
	return fmt.Sprintf("```%s\n%s\n```", b.Code.Language, plainText(b.Code.RichText)), true
}

func renderQuote(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.QuoteBlock)
	if !ok {
		return "", false
	}
	return "> " + plainText(b.Quote.RichText), true
}

func Render(bs []blocks.Block) string {
	renderer := NewRenderer()
	var sb strings.Builder
	for _, b := range bs {
		if line, ok := renderer.Render(b); ok && line != "" {
			sb.WriteString(line)
			sb.WriteString("\n")
		}
	}
	return strings.TrimSpace(sb.String())
}

func plainText(rts []blocks.RichText) string {
	var sb strings.Builder
	for _, rt := range rts {
		sb.WriteString(rt.PlainText)
	}
	return sb.String()
}