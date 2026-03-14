package elements

import (
	"github.com/mathisbot/notionary-go/blocks"
	syntax "github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

type FileCodec struct{}

func (c *FileCodec) Parse(line string) (blocks.Block, bool) {
	syn, ok := syntax.Registry[syntax.File].(syntax.SelfClosingTagSyntax)
	if !ok {
		return nil, false
	}
	m := syn.Pattern.FindStringSubmatch(line)
	if m == nil {
		return nil, false
	}
	return &blocks.FileBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeFile},
		File: blocks.FileData{
			URL:  m[1],
			Name: m[2],
		},
	}, true
}

func (c *FileCodec) Render(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.FileBlock)
	if !ok {
		return "", false
	}
	syn, ok := syntax.Registry[syntax.File].(syntax.SelfClosingTagSyntax)
	if !ok {
		return "", false
	}
	result := syn.Tag + ` src="` + b.File.URL + `"`
	if b.File.Name != "" {
		result += ` name="` + b.File.Name + `"`
	}
	result += ">"
	return result, true
}