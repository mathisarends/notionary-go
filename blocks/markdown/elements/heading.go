package elements

import (
	"github.com/mathisbot/notionary-go/blocks"
	syntax "github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

type HeadingCodec struct{}

func (c *HeadingCodec) Parse(line string) (blocks.Block, bool) {
	syn, ok := syntax.Registry[syntax.Heading].(syntax.SimpleSyntax)
	if !ok {
		return nil, false
	}
	m := syn.Pattern.FindStringSubmatch(line)
	if m == nil {
		return nil, false
	}
	level := len(m[1])
	return &blocks.HeadingBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeHeading},
		Heading: blocks.HeadingData{
			RichText: toRichText(m[2]),
			Level:    level,
			Color:    blocks.BlockColorDefault,
		},
	}, true
}

func (c *HeadingCodec) Render(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.HeadingBlock)
	if !ok {
		return "", false
	}
	syn, ok := syntax.Registry[syntax.Heading].(syntax.SimpleSyntax)
	if !ok {
		return "", false
	}
	prefix := ""
	for i := 0; i < b.Heading.Level; i++ {
		prefix += "#"
	}
	return prefix + " " + toMarkdown(b.Heading.RichText), true
}