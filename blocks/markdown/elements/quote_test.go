package elements

import (
	"testing"

	"github.com/mathisbot/notionary-go/blocks"
)

func TestQuoteCodec_Parse(t *testing.T) {
	c := &QuoteCodec{}
	tests := []struct {
		name  string
		line  string
		valid bool
	}{
		{"simple", `> This is a quote`, true},
		{"double arrow should not match", `>> nested`, false},
		{"no space", `>nospace`, true},
		{"empty", ``, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, ok := c.Parse(tt.line)
			if ok != tt.valid {
				t.Fatalf("expected ok=%v, got %v", tt.valid, ok)
			}
			if tt.valid {
				if _, ok := block.(*blocks.QuoteBlock); !ok {
					t.Fatal("expected *blocks.QuoteBlock")
				}
			}
		})
	}
}

func TestQuoteCodec_Render(t *testing.T) {
	c := &QuoteCodec{}
	tests := []struct {
		name     string
		block    blocks.Block
		expected string
		valid    bool
	}{
		{
			"simple",
			&blocks.QuoteBlock{Quote: blocks.QuoteData{RichText: toRichText("This is a quote")}},
			"> This is a quote",
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