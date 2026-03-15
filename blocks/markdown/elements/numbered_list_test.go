package elements

import (
	"testing"

	"github.com/mathisbot/notionary-go/blocks"
)

func TestNumberedListCodec_Parse(t *testing.T) {
	c := &NumberedListCodec{}
	tests := []struct {
		name  string
		line  string
		valid bool
	}{
		{"simple", `1. First item`, true},
		{"higher number", `42. Some item`, true},
		{"bulleted should not match", `- item`, false},
		{"empty", ``, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, ok := c.Parse(tt.line)
			if ok != tt.valid {
				t.Fatalf("expected ok=%v, got %v", tt.valid, ok)
			}
			if tt.valid {
				if _, ok := block.(*blocks.NumberedListItemBlock); !ok {
					t.Fatal("expected *blocks.NumberedListItemBlock")
				}
			}
		})
	}
}

func TestNumberedListCodec_Render(t *testing.T) {
	c := &NumberedListCodec{}
	tests := []struct {
		name     string
		block    blocks.Block
		expected string
		valid    bool
	}{
		{
			"simple",
			&blocks.NumberedListItemBlock{NumberedListItem: blocks.ListItemData{RichText: toRichText("First item")}},
			"1. First item",
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