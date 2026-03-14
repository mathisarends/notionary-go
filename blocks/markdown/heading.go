package markdown

import (
	"strings"

	"github.com/mathisbot/notionary-go/blocks"
)

type HeadingParser struct {
	BaseParser
}

func (p *HeadingParser) Parse(line string) (any, bool) {
	if text, ok := strings.CutPrefix(line, "### "); ok {
		return blocks.Heading3Block{
			Heading3: blocks.HeadingData{RichText: toRichText(text)},
		}, true
	}
	if text, ok := strings.CutPrefix(line, "## "); ok {
		return blocks.Heading2Block{
			Heading2: blocks.HeadingData{RichText: toRichText(text)},
		}, true
	}
	if text, ok := strings.CutPrefix(line, "# "); ok {
		return blocks.Heading1Block{
			Heading1: blocks.HeadingData{RichText: toRichText(text)},
		}, true
	}
	return p.Next(line)
}