package elements

import (
	"github.com/mathisbot/notionary-go/blocks"
	syntax "github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

type NumberedListCodec struct{}

func (c *NumberedListCodec) Parse(line string) (blocks.Block, bool) {
	syn, ok := syntax.Registry[syntax.NumberedList].(syntax.SimpleSyntax)
	if !ok {
		return nil, false
	}
	m := syn.Pattern.FindStringSubmatch(line)
	if m == nil {
		return nil, false
	}
	return &blocks.NumberedListItemBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeNumberedListItem},
		NumberedListItem: blocks.ListItemData{
			RichText: toRichText(m[3]),
			Color:    blocks.BlockColorDefault,
		},
	}, true
}

func (c *NumberedListCodec) Render(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.NumberedListItemBlock)
	if !ok {
		return "", false
	}
	syn, ok := syntax.Registry[syntax.NumberedList].(syntax.SimpleSyntax)
	if !ok {
		return "", false
	}
	return syn.StartDelimiter + toMarkdown(b.NumberedListItem.RichText), true
}