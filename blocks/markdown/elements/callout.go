package elements

import (
	"github.com/mathisbot/notionary-go/blocks"
	syntax "github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

type CalloutCodec struct{}

func (c *CalloutCodec) Parse(line string) (blocks.Block, bool) {
	syn, ok := syntax.Registry[syntax.Callout].(syntax.TagSyntax)
	if !ok {
		return nil, false
	}
	m := syn.Pattern.FindStringSubmatch(line)
	if m == nil {
		return nil, false
	}
	return &blocks.CalloutBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeCallout},
		Callout: blocks.CalloutData{
			Emoji: m[1],
			Color: blocks.BlockColor(m[2]),
		},
	}, true
}

func (c *CalloutCodec) Render(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.CalloutBlock)
	if !ok {
		return "", false
	}
	syn, ok := syntax.Registry[syntax.Callout].(syntax.TagSyntax)
	if !ok {
		return "", false
	}
	result := syn.OpenTag
	if b.Callout.Emoji != "" {
		result += ` emoji="` + b.Callout.Emoji + `"`
	}
	if b.Callout.Color != "" && b.Callout.Color != blocks.BlockColorDefault {
		result += ` color="` + string(b.Callout.Color) + `"`
	}
	result += ">"
	return result, true
}