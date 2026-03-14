package elements

import (
	"github.com/mathisbot/notionary-go/blocks"
	syntax "github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

type ParagraphCodec struct{}

func (c *ParagraphCodec) Parse(line string) (blocks.Block, bool) {
	return &blocks.ParagraphBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeParagraph},
		Paragraph: blocks.ParagraphData{
			RichText: toRichText(line),
			Color:    blocks.BlockColorDefault,
		},
	}, true
}

func (c *ParagraphCodec) Render(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.ParagraphBlock)
	if !ok {
		return "", false
	}
	_ = syntax.Registry[syntax.Paragraph]
	return toMarkdown(b.Paragraph.RichText), true
}