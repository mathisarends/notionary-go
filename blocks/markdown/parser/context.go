package parser

import "github.com/mathisbot/notionary-go/blocks"

const spacesPerLevel = 2

type ParentBlockContext struct {
	Block       blocks.Block
	ChildLines  []string
	ChildBlocks []blocks.Block
}

func (p *ParentBlockContext) AddChildLine(line string) {
	p.ChildLines = append(p.ChildLines, line)
}

func (p *ParentBlockContext) AddChildBlock(block blocks.Block) {
	p.ChildBlocks = append(p.ChildBlocks, block)
}

type ParseChildrenCallback func(text string) ([]blocks.Block, error)

type Context struct {
	Line                string
	ResultBlocks        *[]blocks.Block
	ParentStack         *[]ParentBlockContext
	AllLines            []string
	CurrentLineIndex    int
	LinesConsumed       int
	IsPreviousLineEmpty bool
	ParseChildren       ParseChildrenCallback
}

func (c *Context) GetRemainingLines() []string {
	if c.CurrentLineIndex+1 >= len(c.AllLines) {
		return nil
	}
	return c.AllLines[c.CurrentLineIndex+1:]
}

func (c *Context) IsInsideParentContext() bool {
	return len(*c.ParentStack) > 0
}

func (c *Context) GetIndentLevel(line string) int {
	spaces := len(line) - len(trimLeft(line))
	return spaces / spacesPerLevel
}

func (c *Context) CollectIndentedChildLines(parentIndent int) []string {
	expected := parentIndent + 1
	var children []string
	for _, line := range c.GetRemainingLines() {
		if line == "" || len(trimLeft(line)) == 0 {
			children = append(children, line)
			continue
		}
		if c.GetIndentLevel(line) >= expected {
			children = append(children, line)
		} else {
			break
		}
	}
	return children
}

func (c *Context) StripIndent(lines []string, levels int) []string {
	spaces := spacesPerLevel * levels
	result := make([]string, len(lines))
	for i, line := range lines {
		if len(line) == 0 || len(trimLeft(line)) == 0 {
			result[i] = line
			continue
		}
		if len(line) < spaces {
			result[i] = line
		} else {
			result[i] = line[spaces:]
		}
	}
	return result
}

func trimLeft(s string) string {
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	return s[i:]
}