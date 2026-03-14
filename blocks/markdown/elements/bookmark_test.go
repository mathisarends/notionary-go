package elements

import (
	"testing"

	"github.com/mathisbot/notionary-go/blocks"
)

func TestBookmarkCodec_Parse(t *testing.T) {
	c := &BookmarkCodec{}
	tests := []struct {
		name  string
		line  string
		valid bool
	}{
		{"with title", `<bookmark url="https://example.com" title="My Site">`, true},
		{"without title", `<bookmark url="https://example.com">`, true},
		{"invalid", `bookmark url="https://example.com"`, false},
		{"empty", ``, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, ok := c.Parse(tt.line)
			if ok != tt.valid {
				t.Fatalf("expected ok=%v, got %v", tt.valid, ok)
			}
			if !tt.valid {
				return
			}
			b, ok := block.(*blocks.BookmarkBlock)
			if !ok {
				t.Fatal("expected *blocks.BookmarkBlock")
			}
			if b.Bookmark.URL == "" {
				t.Error("expected non-empty URL")
			}
		})
	}
}

func TestBookmarkCodec_Render(t *testing.T) {
	c := &BookmarkCodec{}
	tests := []struct {
		name     string
		block    blocks.Block
		expected string
		valid    bool
	}{
		{
			"with title",
			&blocks.BookmarkBlock{Bookmark: blocks.BookmarkData{URL: "https://example.com", Title: toRichText("My Site")}},
			`<bookmark url="https://example.com" title="My Site">`,
			true,
		},
		{
			"without title",
			&blocks.BookmarkBlock{Bookmark: blocks.BookmarkData{URL: "https://example.com"}},
			`<bookmark url="https://example.com">`,
			true,
		},
		{"wrong type", &blocks.ParagraphBlock{}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := c.Render(tt.block)
			if ok != tt.valid {
				t.Fatalf("expected ok=%v, got %v", tt.valid, ok)
			}
			if tt.valid && got != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}