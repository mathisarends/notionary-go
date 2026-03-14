package elements

import (
	"github.com/mathisbot/notionary-go/blocks"
	syntax "github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

type BreadcrumbCodec struct{}

func (c *BreadcrumbCodec) Parse(line string) (blocks.Block, bool) {
	syn, ok := syntax.Registry[syntax.Breadcrumb].(syntax.SimpleSyntax)
	if !ok {
		return nil, false
	}
	if !syn.Pattern.MatchString(line) {
		return nil, false
	}
	return &blocks.BreadcrumbBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeBreadcrumb},
	}, true
}

func (c *BreadcrumbCodec) Render(block blocks.Block) (string, bool) {
	_, ok := block.(*blocks.BreadcrumbBlock)
	if !ok {
		return "", false
	}
	syn, ok := syntax.Registry[syntax.Breadcrumb].(syntax.SimpleSyntax)
	if !ok {
		return "", false
	}
	return syn.StartDelimiter, true
}