package elements

import (
	"github.com/mathisbot/notionary-go/blocks"
	syntax "github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

type ColumnListCodec struct{}

func (c *ColumnListCodec) Parse(line string) (blocks.Block, bool) {
	syn, ok := syntax.Registry[syntax.ColumnList].(syntax.TagSyntax)
	if !ok {
		return nil, false
	}
	if !syn.Pattern.MatchString(line) {
		return nil, false
	}
	return &blocks.ColumnListBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeColumnList},
	}, true
}

func (c *ColumnListCodec) Render(block blocks.Block) (string, bool) {
	_, ok := block.(*blocks.ColumnListBlock)
	if !ok {
		return "", false
	}
	syn, ok := syntax.Registry[syntax.ColumnList].(syntax.TagSyntax)
	if !ok {
		return "", false
	}
	return syn.OpenTag + ">", true
}