package elements

import (
	"testing"

	"github.com/mathisbot/notionary-go/blocks"
)

func TestBulletedListCodec_Parse(t *testing.T) {
	c := &BulletedListCodec{}
	tests := []struct {
		name  string
		line  string
		valid bool
	}{
		{"simple", `- First item`, true},
		{"nested", `  - Nested item`, true},
		{"todo should not match", `- [ ] Task`, false},
		{"numbered should not match", `1. Item`, false},
		{"empty", ``, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, ok := c.Parse(tt.line)
			if ok != tt.valid {
				t.Fatalf("expected ok=%v, got %v", tt.valid, ok)
			}
			if tt.valid {
				if _, ok := block.(*blocks.BulletedListItemBlock); !ok {
					t.Fatal("expected *blocks.BulletedListItemBlock")
				}
			}
		})
	}
}

func TestBulletedListCodec_Render(t *testing.T) {
	c := &BulletedListCodec{}
	tests := []struct {
		name     string
		block    blocks.Block
		expected string
		valid    bool
	}{
		{
			"simple",
			&blocks.BulletedListItemBlock{BulletedListItem: blocks.ListItemData{RichText: toRichText("First item")}},
			"- First item",
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