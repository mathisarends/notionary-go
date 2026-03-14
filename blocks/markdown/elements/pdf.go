package elements

import (
	"github.com/mathisbot/notionary-go/blocks"
	syntax "github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

type PDFCodec struct{}

func (c *PDFCodec) Parse(line string) (blocks.Block, bool) {
	syn, ok := syntax.Registry[syntax.PDF].(syntax.SelfClosingTagSyntax)
	if !ok {
		return nil, false
	}
	m := syn.Pattern.FindStringSubmatch(line)
	if m == nil {
		return nil, false
	}
	return &blocks.PDFBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypePDF},
		PDF: blocks.FileData{
			URL:     m[1],
			Caption: toRichText(m[2]),
		},
	}, true
}

func (c *PDFCodec) Render(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.PDFBlock)
	if !ok {
		return "", false
	}
	syn, ok := syntax.Registry[syntax.PDF].(syntax.SelfClosingTagSyntax)
	if !ok {
		return "", false
	}
	result := syn.Tag + ` src="` + b.PDF.URL + `"`
	if caption := toMarkdown(b.PDF.Caption); caption != "" {
		result += ` caption="` + caption + `"`
	}
	result += ">"
	return result, true
}