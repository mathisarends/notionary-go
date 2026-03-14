package elements

import (
	"github.com/mathisbot/notionary-go/blocks"
	syntax "github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

type EquationCodec struct{}

func (c *EquationCodec) Parse(line string) (blocks.Block, bool) {
	syn, ok := syntax.Registry[syntax.Equation].(syntax.TagSyntax)
	if !ok {
		return nil, false
	}
	if !syn.Pattern.MatchString(line) {
		return nil, false
	}
	return &blocks.EquationBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeEquation},
	}, true
}

func (c *EquationCodec) Render(block blocks.Block) (string, bool) {
	_, ok := block.(*blocks.EquationBlock)
	if !ok {
		return "", false
	}
	syn, ok := syntax.Registry[syntax.Equation].(syntax.TagSyntax)
	if !ok {
		return "", false
	}
	return syn.OpenTag + ">", true
}