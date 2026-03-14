package elements

import (
	"github.com/mathisbot/notionary-go/blocks"
	syntax "github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

type ColumnListCodec struct{}
type ColumnCodec struct{}

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
	return syn.OpenTag, true
}

func (c *ColumnCodec) Parse(line string) (blocks.Block, bool) {
	syn, ok := syntax.Registry[syntax.Column].(syntax.TagSyntax)
	if !ok {
		return nil, false
	}
	m := syn.Pattern.FindStringSubmatch(line)
	if m == nil {
		return nil, false
	}
	return &blocks.ColumnBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeColumn},
		Column: blocks.ColumnData{
			Ratio: m[1],
		},
	}, true
}

func (c *ColumnCodec) Render(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.ColumnBlock)
	if !ok {
		return "", false
	}
	syn, ok := syntax.Registry[syntax.Column].(syntax.TagSyntax)
	if !ok {
		return "", false
	}
	result := syn.OpenTag
	if b.Column.Ratio != "" {
		result += ` ratio="` + b.Column.Ratio + `"`
	}
	result += ">"
	return result, true
}