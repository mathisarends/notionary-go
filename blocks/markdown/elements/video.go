package elements

import (
	"github.com/mathisbot/notionary-go/blocks"
	syntax "github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

type VideoCodec struct{}

func (c *VideoCodec) Parse(line string) (blocks.Block, bool) {
	syn, ok := syntax.Registry[syntax.Video].(syntax.SelfClosingTagSyntax)
	if !ok {
		return nil, false
	}
	m := syn.Pattern.FindStringSubmatch(line)
	if m == nil {
		return nil, false
	}
	return &blocks.VideoBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeVideo},
		Video: blocks.FileData{
			URL:     m[1],
			Caption: toRichText(m[2]),
		},
	}, true
}

func (c *VideoCodec) Render(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.VideoBlock)
	if !ok {
		return "", false
	}
	syn, ok := syntax.Registry[syntax.Video].(syntax.SelfClosingTagSyntax)
	if !ok {
		return "", false
	}
	result := syn.Tag + ` src="` + b.Video.URL + `"`
	if caption := toMarkdown(b.Video.Caption); caption != "" {
		result += ` caption="` + caption + `"`
	}
	result += ">"
	return result, true
}