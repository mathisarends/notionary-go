package elements

import (
	"github.com/mathisbot/notionary-go/blocks"
	syntax "github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

type BookmarkCodec struct{}

func (c *BookmarkCodec) Parse(line string) (blocks.Block, bool) {
	syn, ok := syntax.Registry[syntax.Bookmark].(syntax.SelfClosingTagSyntax)
	if !ok {
		return nil, false
	}
	m := syn.Pattern.FindStringSubmatch(line)
	if m == nil {
		return nil, false
	}
	return &blocks.BookmarkBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeBookmark},
		Bookmark: blocks.BookmarkData{
			URL:   m[1],
			Title: toRichText(m[2]),
		},
	}, true
}

func (c *BookmarkCodec) Render(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.BookmarkBlock)
	if !ok {
		return "", false
	}
	syn, ok := syntax.Registry[syntax.Bookmark].(syntax.SelfClosingTagSyntax)
	if !ok {
		return "", false
	}
	result := syn.Tag + ` url="` + b.Bookmark.URL + `"`
	if title := toMarkdown(b.Bookmark.Title); title != "" {
		result += ` title="` + title + `"`
	}
	result += ">"
	return result, true
}