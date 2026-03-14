package elements

import (
	"github.com/mathisbot/notionary-go/blocks"
	syntax "github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

type ToDoCodec struct{}
type ToDoDoneCodec struct{}

func (c *ToDoCodec) Parse(line string) (blocks.Block, bool) {
	syn, ok := syntax.Registry[syntax.ToDo].(syntax.SimpleSyntax)
	if !ok {
		return nil, false
	}
	m := syn.Pattern.FindStringSubmatch(line)
	if m == nil {
		return nil, false
	}
	return &blocks.ToDoBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeToDo},
		ToDo: blocks.ToDoData{
			RichText: toRichText(m[1]),
			Checked:  false,
			Color:    blocks.BlockColorDefault,
		},
	}, true
}

func (c *ToDoCodec) Render(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.ToDoBlock)
	if !ok || b.ToDo.Checked {
		return "", false
	}
	syn, ok := syntax.Registry[syntax.ToDo].(syntax.SimpleSyntax)
	if !ok {
		return "", false
	}
	return syn.StartDelimiter + " " + toMarkdown(b.ToDo.RichText), true
}


func (c *ToDoDoneCodec) Parse(line string) (blocks.Block, bool) {
	syn, ok := syntax.Registry[syntax.ToDoDone].(syntax.SimpleSyntax)
	if !ok {
		return nil, false
	}
	m := syn.Pattern.FindStringSubmatch(line)
	if m == nil {
		return nil, false
	}
	return &blocks.ToDoBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeToDo},
		ToDo: blocks.ToDoData{
			RichText: toRichText(m[1]),
			Checked:  true,
			Color:    blocks.BlockColorDefault,
		},
	}, true
}

func (c *ToDoDoneCodec) Render(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.ToDoBlock)
	if !ok || !b.ToDo.Checked {
		return "", false
	}
	syn, ok := syntax.Registry[syntax.ToDoDone].(syntax.SimpleSyntax)
	if !ok {
		return "", false
	}
	return syn.StartDelimiter + " " + toMarkdown(b.ToDo.RichText), true
}