package markdown

import (
	"regexp"

	"github.com/mathisbot/notionary-go/blocks"
)

var bulletedListPattern = regexp.MustCompile(`^(\s*)[-*+]\s+(?!\[[ xX]\])(.+)$`)

type BulletedListParser struct {
	BaseParser
	pattern *regexp.Regexp
}

type BulletedListRenderer struct {
}

func (p *BulletedListParser) Parse(line string) (any, bool) {
	m := p.pattern.FindStringSubmatch(line)
	if m == nil {
		return p.Next(line)
	}

	return blocks.BulletedListItemBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeBulletedListItem},
		BulletedListItem: blocks.ListItemData{
			RichText: toRichText(m[2]),
			Color:    blocks.BlockColorDefault,
		},
	}, true
}

func NewBulletedListParser(pattern *regexp.Regexp) *BulletedListParser {
	if pattern == nil {
		pattern = bulletedListPattern
	}
	return &BulletedListParser{pattern: pattern}
}

func renderBulletedListItem(block *blocks.BulletedListItemBlock) string {
	return "- " + plainText(block.BulletedListItem.RichText)
}

func (r *BulletedListRenderer) Render(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.BulletedListItemBlock)
	if !ok {
		return "", false
	}
	return renderBulletedListItem(b), true
}
