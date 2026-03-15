package elements

import (
	"github.com/mathisbot/notionary-go/blocks"
	syntax "github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

type CodeCodec struct{}

// TODO: Sollte auf ``` markdown syntax ausgetauscht werden
func (c *CodeCodec) Parse(line string) (blocks.Block, bool) {
	syn, ok := syntax.Registry[syntax.Code].(syntax.TagSyntax)
	if !ok {
		return nil, false
	}
	m := syn.Pattern.FindStringSubmatch(line)
	if m == nil {
		return nil, false
	}
	return &blocks.CodeBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeCode},
		Code: blocks.CodeData{
			Language: m[1],
		},
	}, true
}

func (c *CodeCodec) Render(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.CodeBlock)
	if !ok {
		return "", false
	}
	syn, ok := syntax.Registry[syntax.Code].(syntax.TagSyntax)
	if !ok {
		return "", false
	}
	result := syn.OpenTag
	if b.Code.Language != "" {
		result += ` lang="` + b.Code.Language + `"`
	}
	result += ">"
	return result, true
}