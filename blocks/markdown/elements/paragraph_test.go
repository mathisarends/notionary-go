package elements

import (
	"testing"

	"github.com/mathisbot/notionary-go/blocks"
)

func TestParagraphCodec_Parse(t *testing.T) {
	c := &ParagraphCodec{}
	tests := []struct {
		name  string
		line  string
	}{
		{"normal text", `Hello World`},
		{"empty line", ``},
		{"special chars", `foo & bar <baz>`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, ok := c.Parse(tt.line)
			if !ok {
				t.Fatal("ParagraphCodec.Parse should always return true")
			}
			if _, ok := block.(*blocks.ParagraphBlock); !ok {
				t.Fatal("expected *blocks.ParagraphBlock")
			}
		})
	}
}

func TestParagraphCodec_Render(t *testing.T) {
	c := &ParagraphCodec{}
	tests := []struct {
		name     string
		block    blocks.Block
		expected string
		valid    bool
	}{
		{
			"simple",
			&blocks.ParagraphBlock{Paragraph: blocks.ParagraphData{RichText: toRichText("Hello World")}},
			"Hello World",
			true,
		},
		{"wrong type", &blocks.DividerBlock{}, "", false},
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