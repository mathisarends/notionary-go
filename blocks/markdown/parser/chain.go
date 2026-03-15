package parser

import "github.com/mathisbot/notionary-go/blocks"

type LineParser interface {
	Parse(line string) (blocks.Block, bool)
}

type chain struct {
	parsers []LineParser
}

func newChain(parsers ...LineParser) *chain {
	return &chain{parsers: parsers}
}

func (c *chain) Parse(line string) (blocks.Block, bool) {
	for _, p := range c.parsers {
		if block, ok := p.Parse(line); ok {
			return block, true
		}
	}
	return nil, false
}