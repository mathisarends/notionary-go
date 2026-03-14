package elements

import (
	"github.com/mathisbot/notionary-go/blocks"
	syntax "github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

type TableCodec struct{}
type TableRowCodec struct{}

func (c *TableCodec) Parse(line string) (blocks.Block, bool) {
	syn, ok := syntax.Registry[syntax.Table].(syntax.SimpleSyntax)
	if !ok {
		return nil, false
	}
	if !syn.Pattern.MatchString(line) {
		return nil, false
	}
	return &blocks.TableBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeTable},
	}, true
}

func (c *TableCodec) Render(block blocks.Block) (string, bool) {
	_, ok := block.(*blocks.TableBlock)
	if !ok {
		return "", false
	}
	syn, ok := syntax.Registry[syntax.Table].(syntax.SimpleSyntax)
	if !ok {
		return "", false
	}
	return syn.StartDelimiter, true
}

func (c *TableRowCodec) Parse(line string) (blocks.Block, bool) {
	syn, ok := syntax.Registry[syntax.TableRow].(syntax.SimpleSyntax)
	if !ok {
		return nil, false
	}
	if !syn.Pattern.MatchString(line) {
		return nil, false
	}
	return &blocks.TableRowBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeTableRow},
	}, true
}

func (c *TableRowCodec) Render(block blocks.Block) (string, bool) {
	_, ok := block.(*blocks.TableRowBlock)
	if !ok {
		return "", false
	}
	syn, ok := syntax.Registry[syntax.TableRow].(syntax.SimpleSyntax)
	if !ok {
		return "", false
	}
	return syn.StartDelimiter, true
}