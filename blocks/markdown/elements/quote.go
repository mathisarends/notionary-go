package elements

import (
	"strings"

	"github.com/mathisbot/notionary-go/blocks"
	syntax "github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

type QuoteCodec struct{}

func (c *QuoteCodec) Parse(line string) (blocks.Block, bool) {
	if strings.HasPrefix(strings.TrimSpace(line), ">>") {
		return nil, false
	}

	syn, ok := syntax.Registry[syntax.Quote].(syntax.SimpleSyntax)
	if !ok {
		return nil, false
	}
	m := syn.Pattern.FindStringSubmatch(line)
	if m == nil {
		return nil, false
	}
	return &blocks.QuoteBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeQuote},
		Quote: blocks.QuoteData{
			RichText: toRichText(m[1]),
			Color:    blocks.BlockColorDefault,
		},
	}, true
}

func (c *QuoteCodec) Render(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.QuoteBlock)
	if !ok {
		return "", false
	}
	syn, ok := syntax.Registry[syntax.Quote].(syntax.SimpleSyntax)
	if !ok {
		return "", false
	}
	return syn.StartDelimiter + toMarkdown(b.Quote.RichText), true
}