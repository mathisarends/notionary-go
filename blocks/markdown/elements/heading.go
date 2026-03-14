package elements

import (
	"github.com/mathisbot/notionary-go/blocks"
	syn "github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

type HeadingCodec struct{}

func (c *HeadingCodec) Parse(line string) (blocks.Block, bool) {
	syntax, ok := syn.Registry[syn.Heading].(syn.SimpleSyntax)
	if !ok {
		return nil, false
	}
	m := syntax.Pattern.FindStringSubmatch(line)
	if m == nil {
		return nil, false
	}
	level := len(m[1])
	data := blocks.HeadingData{
		RichText: toRichText(m[2]),
		Color:    blocks.BlockColorDefault,
	}
	switch level {
	case 1:
		return &blocks.Heading1Block{BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeHeading1}, Heading1: data}, true
	case 2:
		return &blocks.Heading2Block{BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeHeading2}, Heading2: data}, true
	case 3:
		return &blocks.Heading3Block{BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeHeading3}, Heading3: data}, true
	default:
		return nil, false
	}
}

func (c *HeadingCodec) Render(block blocks.Block) (string, bool) {
	syntax, ok := syn.Registry[syn.Heading].(syn.SimpleSyntax)
	if !ok {
		return "", false
	}
	_ = syntax
	var level int
	var richText []blocks.RichText
	switch b := block.(type) {
	case *blocks.Heading1Block:
		level, richText = 1, b.Heading1.RichText
	case *blocks.Heading2Block:
		level, richText = 2, b.Heading2.RichText
	case *blocks.Heading3Block:
		level, richText = 3, b.Heading3.RichText
	default:
		return "", false
	}
	prefix := ""
	for i := 0; i < level; i++ {
		prefix += "#"
	}
	return prefix + " " + toMarkdown(richText), true
}