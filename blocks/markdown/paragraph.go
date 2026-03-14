package markdown

import (
	"github.com/mathisbot/notionary-go/blocks"
)

// ParagraphParser ist immer letztes Glied — kein Next
type ParagraphParser struct{}

func (p *ParagraphParser) Parse(line string) (any, bool) {
	if line == "" {
		return nil, false
	}
	return blocks.ParagraphBlock{
		Paragraph: blocks.ParagraphData{
			RichText: toRichText(line),
		},
	}, true
}

func (p *ParagraphParser) SetNext(next Parser) Parser { return next }