package markdown

import (
	"regexp"

	"github.com/mathisbot/notionary-go/blocks"
)

const numberedListPlaceholder = "__NUM__"

var numberedListPattern = regexp.MustCompile(`^(\s*)(\d+)\.\s+(.+)$`)

type NumberedListParser struct {
	BaseParser
	pattern *regexp.Regexp
}

type NumberedListRenderer struct {
}

func (p *NumberedListParser) Parse(line string) (any, bool) {
	m := p.pattern.FindStringSubmatch(line)
	if m == nil {
		return p.Next(line)
	}

	return blocks.NumberedListItemBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeNumberedListItem},
		NumberedListItem: blocks.ListItemData{
			RichText: toRichText(m[3]),
			Color:    blocks.BlockColorDefault,
		},
	}, true
}

func NewNumberedListParser(pattern *regexp.Regexp) *NumberedListParser {
	if pattern == nil {
		pattern = numberedListPattern
	}
	return &NumberedListParser{pattern: pattern}
}

func renderNumberedListItem(block *blocks.NumberedListItemBlock) string {
	return numberedListPlaceholder + ". " + plainText(block.NumberedListItem.RichText)
}

func (r *NumberedListRenderer) Render(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.NumberedListItemBlock)
	if !ok {
		return "", false
	}
	return renderNumberedListItem(b), true
}
