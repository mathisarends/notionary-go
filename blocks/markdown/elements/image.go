package elements

import (
	"github.com/mathisbot/notionary-go/blocks"
	syntax "github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

type ImageCodec struct{}

func (c *ImageCodec) Parse(line string) (blocks.Block, bool) {
	syn, ok := syntax.Registry[syntax.Image].(syntax.SelfClosingTagSyntax)
	if !ok {
		return nil, false
	}
	m := syn.Pattern.FindStringSubmatch(line)
	if m == nil {
		return nil, false
	}
	return &blocks.ImageBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeImage},
		Image: blocks.FileData{
			URL:     m[1],
			Caption: toRichText(m[2]),
		},
	}, true
}

func (c *ImageCodec) Render(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.ImageBlock)
	if !ok {
		return "", false
	}
	syn, ok := syntax.Registry[syntax.Image].(syntax.SelfClosingTagSyntax)
	if !ok {
		return "", false
	}
	result := syn.Tag + ` src="` + b.Image.URL + `"`
	if caption := toMarkdown(b.Image.Caption); caption != "" {
		result += ` caption="` + caption + `"`
	}
	result += ">"
	return result, true
}