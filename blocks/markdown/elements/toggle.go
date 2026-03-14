package elements

import (
	"github.com/mathisbot/notionary-go/blocks"
	syntax "github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

type ToggleCodec struct{}

func (c *ToggleCodec) Parse(line string) (blocks.Block, bool) {
	syn, ok := syntax.Registry[syntax.Toggle].(syntax.TagSyntax)
	if !ok {
		return nil, false
	}
	m := syn.Pattern.FindStringSubmatch(line)
	if m == nil {
		return nil, false
	}
	return &blocks.ToggleBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeToggle},
		Toggle: blocks.ToggleData{
			Title: toRichText(m[1]),
		},
	}, true
}

func (c *ToggleCodec) Render(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.ToggleBlock)
	if !ok {
		return "", false
	}
	syn, ok := syntax.Registry[syntax.Toggle].(syntax.TagSyntax)
	if !ok {
		return "", false
	}
	result := syn.OpenTag
	if title := toMarkdown(b.Toggle.Title); title != "" {
		result += ` title="` + title + `"`
	}
	result += ">"
	return result, true
}