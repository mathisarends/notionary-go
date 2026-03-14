package elements

import (
	"github.com/mathisbot/notionary-go/blocks"
	syntax "github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

type AudioCodec struct{}

func (c *AudioCodec) Parse(line string) (blocks.Block, bool) {
	syn, ok := syntax.Registry[syntax.Audio].(syntax.SelfClosingTagSyntax)
	if !ok {
		return nil, false
	}
	m := syn.Pattern.FindStringSubmatch(line)
	if m == nil {
		return nil, false
	}
	return &blocks.AudioBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeAudio},
		Audio: blocks.FileData{
			URL:     m[1],
			Caption: toRichText(m[2]),
		},
	}, true
}

func (c *AudioCodec) Render(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.AudioBlock)
	if !ok {
		return "", false
	}
	syn, ok := syntax.Registry[syntax.Audio].(syntax.SelfClosingTagSyntax)
	if !ok {
		return "", false
	}
	result := syn.Tag + ` src="` + b.Audio.URL + `"`
	if caption := toMarkdown(b.Audio.Caption); caption != "" {
		result += ` caption="` + caption + `"`
	}
	result += ">"
	return result, true
}