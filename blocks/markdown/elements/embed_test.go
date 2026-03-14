package elements

import (
	"testing"

	"github.com/mathisbot/notionary-go/blocks"
)

func TestEmbedCodec_Parse(t *testing.T) {
	c := &EmbedCodec{}
	tests := []struct {
		name  string
		line  string
		valid bool
	}{
		{"with title", `<embed url="https://example.com" title="My Embed">`, true},
		{"without title", `<embed url="https://example.com">`, true},
		{"invalid", `embed url="https://example.com"`, false},
		{"empty", ``, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, ok := c.Parse(tt.line)
			if ok != tt.valid {
				t.Fatalf("expected ok=%v, got %v", tt.valid, ok)
			}
			if tt.valid {
				if _, ok := block.(*blocks.EmbedBlock); !ok {
					t.Fatal("expected *blocks.EmbedBlock")
				}
			}
		})
	}
}

func TestEmbedCodec_Render(t *testing.T) {
	c := &EmbedCodec{}
	tests := []struct {
		name     string
		block    blocks.Block
		expected string
		valid    bool
	}{
		{
			"with title",
			&blocks.EmbedBlock{Embed: blocks.EmbedData{URL: "https://example.com", Title: toRichText("My Embed")}},
			`<embed url="https://example.com" title="My Embed">`,
			true,
		},
		{
			"without title",
			&blocks.EmbedBlock{Embed: blocks.EmbedData{URL: "https://example.com"}},
			`<embed url="https://example.com">`,
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