package elements

import (
	"strconv"

	"github.com/mathisbot/notionary-go/blocks"
	syntax "github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

type ToggleableHeadingCodec struct{}

func (c *ToggleableHeadingCodec) Parse(line string) (blocks.Block, bool) {
	syn, ok := syntax.Registry[syntax.ToggleableHeading].(syntax.TagSyntax)
	if !ok {
		return nil, false
	}
	m := syn.Pattern.FindStringSubmatch(line)
	if m == nil {
		return nil, false
	}
	level, err := strconv.Atoi(m[2])
	if err != nil {
		return nil, false
	}
	return &blocks.ToggleableHeadingBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeToggleableHeading},
		ToggleableHeading: blocks.ToggleableHeadingData{
			Title: toRichText(m[1]),
			Level: level,
		},
	}, true
}

func (c *ToggleableHeadingCodec) Render(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.ToggleableHeadingBlock)
	if !ok {
		return "", false
	}
	syn, ok := syntax.Registry[syntax.ToggleableHeading].(syntax.TagSyntax)
	if !ok {
		return "", false
	}
	result := syn.OpenTag
	if title := toMarkdown(b.ToggleableHeading.Title); title != "" {
		result += ` title="` + title + `"`
	}
	result += ` level="` + strconv.Itoa(b.ToggleableHeading.Level) + `"`
	result += ">"
	return result, true
}