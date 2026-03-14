package elements

import (
	"github.com/mathisbot/notionary-go/blocks"
	syntax "github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

type DividerCodec struct{}

func (c *DividerCodec) Parse(line string) (blocks.Block, bool) {
	syn, ok := syntax.Registry[syntax.Divider].(syntax.SimpleSyntax)
	if !ok {
		return nil, false
	}
	if !syn.Pattern.MatchString(line) {
		return nil, false
	}
	return &blocks.DividerBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeDivider},
	}, true
}

func (c *DividerCodec) Render(block blocks.Block) (string, bool) {
	_, ok := block.(*blocks.DividerBlock)
	if !ok {
		return "", false
	}
	syn, ok := syntax.Registry[syntax.Divider].(syntax.SimpleSyntax)
	if !ok {
		return "", false
	}
	return syn.StartDelimiter, true
}