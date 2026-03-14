package elements

import (
	"github.com/mathisbot/notionary-go/blocks"
	syntax "github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

type EmbedCodec struct{}

func (c *EmbedCodec) Parse(line string) (blocks.Block, bool) {
	syn, ok := syntax.Registry[syntax.Embed].(syntax.SelfClosingTagSyntax)
	if !ok {
		return nil, false
	}
	m := syn.Pattern.FindStringSubmatch(line)
	if m == nil {
		return nil, false
	}
	return &blocks.EmbedBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeEmbed},
		Embed: blocks.EmbedData{
			URL:   m[1],
			Title: toRichText(m[2]),
		},
	}, true
}

func (c *EmbedCodec) Render(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.EmbedBlock)
	if !ok {
		return "", false
	}
	syn, ok := syntax.Registry[syntax.Embed].(syntax.SelfClosingTagSyntax)
	if !ok {
		return "", false
	}
	result := syn.Tag + ` url="` + b.Embed.URL + `"`
	if title := toMarkdown(b.Embed.Title); title != "" {
		result += ` title="` + title + `"`
	}
	result += ">"
	return result, true
}