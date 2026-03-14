package markdown

import (
	"strings"

	"github.com/mathisbot/notionary-go/blocks"
)

type BulletedListParser struct {
	BaseParser
}

func (p *BulletedListParser) Parse(line string) (any, bool) {
	for _, prefix := range []string{"- ", "* ", "+ "} {
		if strings.HasPrefix(line, prefix) {
			return blocks.BulletedListItemBlock{
				BulletedListItem: blocks.ListItemData{
					RichText: toRichText(strings.TrimPrefix(line, prefix)),
				},
			}, true
		}
	}
	return p.Next(line)
}