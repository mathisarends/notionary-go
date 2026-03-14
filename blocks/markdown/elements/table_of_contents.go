package elements

import (
	"github.com/mathisbot/notionary-go/blocks"
	syntax "github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

type TableOfContentsCodec struct{}

func (c *TableOfContentsCodec) Parse(line string) (blocks.Block, bool) {
	syn, ok := syntax.Registry[syntax.TableOfContents].(syntax.SimpleSyntax)
	if !ok {
		return nil, false
	}
	if !syn.Pattern.MatchString(line) {
		return nil, false
	}
	return &blocks.TableOfContentsBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeTableOfContents},
	}, true
}

func (c *TableOfContentsCodec) Render(block blocks.Block) (string, bool) {
	_, ok := block.(*blocks.TableOfContentsBlock)
	if !ok {
		return "", false
	}
	syn, ok := syntax.Registry[syntax.TableOfContents].(syntax.SimpleSyntax)
	if !ok {
		return "", false
	}
	return syn.StartDelimiter, true
}