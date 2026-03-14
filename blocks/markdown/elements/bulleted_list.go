package elements

import (
    "strings"

	"github.com/mathisbot/notionary-go/blocks"
	syntax "github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

type BulletedListCodec struct{}

func (c *BulletedListCodec) Parse(line string) (blocks.Block, bool) {
    trimmed := strings.TrimSpace(line)
    if strings.HasPrefix(trimmed, "- [ ]") || strings.HasPrefix(strings.ToLower(trimmed), "- [x]") {
        return nil, false
    }

    syntax, ok := syntax.Registry[syntax.BulletedList].(syntax.SimpleSyntax)
    if !ok {
        return nil, false
    }
    m := syntax.Pattern.FindStringSubmatch(line)
    if m == nil {
        return nil, false
    }
    return &blocks.BulletedListItemBlock{
        BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeBulletedListItem},
        BulletedListItem: blocks.ListItemData{
            RichText: toRichText(m[2]),
            Color:    blocks.BlockColorDefault,
        },
    }, true
}

func (c *BulletedListCodec) Render(block blocks.Block) (string, bool) {
    b, ok := block.(*blocks.BulletedListItemBlock)
    if !ok {
        return "", false
    }
    syntax, ok := syntax.Registry[syntax.BulletedList].(syntax.SimpleSyntax)
    if !ok {
        return "", false
    }
    return syntax.StartDelimiter + toMarkdown(b.BulletedListItem.RichText), true
}