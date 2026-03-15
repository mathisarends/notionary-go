package renderer

import (
	"strings"

	"github.com/mathisbot/notionary-go/blocks"
)

type ConvertChildren func(blocks []blocks.Block, indentLevel int) (string, error)

type Context struct {
	Block          blocks.Block
	IndentLevel    int
	convertChildren ConvertChildren

	Result string
}

func newContext(block blocks.Block, indentLevel int, fn ConvertChildren) *Context {
	return &Context{
		Block:           block,
		IndentLevel:     indentLevel,
		convertChildren: fn,
	}
}

func (c *Context) RenderChildren() (string, error) {
	return c.renderChildrenWithIndent(c.IndentLevel)
}

func (c *Context) RenderChildrenIndented(extra int) (string, error) {
	return c.renderChildrenWithIndent(c.IndentLevel + extra)
}

func (c *Context) renderChildrenWithIndent(level int) (string, error) {
	b, ok := c.Block.(blocks.BlockWithChildren)
	if !ok || c.convertChildren == nil {
		return "", nil
	}
	children := b.Children()
	if len(children) == 0 {
		return "", nil
	}
	return c.convertChildren(children, level)
}

const spacesPerLevel = 2

func (c *Context) IndentText(text string) string {
	if text == "" || c.IndentLevel == 0 {
		return text
	}
	prefix := strings.Repeat(" ", spacesPerLevel*c.IndentLevel)
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		if strings.TrimSpace(line) != "" {
			lines[i] = prefix + line
		}
	}
	return strings.Join(lines, "\n")
}