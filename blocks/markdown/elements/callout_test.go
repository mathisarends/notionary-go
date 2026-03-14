package elements

import (
	"testing"

	"github.com/mathisbot/notionary-go/blocks"
)

func TestCalloutCodec_Parse(t *testing.T) {
	c := &CalloutCodec{}
	tests := []struct {
		name  string
		line  string
		valid bool
	}{
		{"with emoji and color", `<callout emoji="💡" color="blue">`, true},
		{"with emoji only", `<callout emoji="💡">`, true},
		{"no attributes", `<callout>`, true},
		{"invalid", `callout emoji="💡"`, false},
		{"empty", ``, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, ok := c.Parse(tt.line)
			if ok != tt.valid {
				t.Fatalf("expected ok=%v, got %v", tt.valid, ok)
			}
			if tt.valid {
				if _, ok := block.(*blocks.CalloutBlock); !ok {
					t.Fatal("expected *blocks.CalloutBlock")
				}
			}
		})
	}
}

func TestCalloutCodec_Render(t *testing.T) {
	c := &CalloutCodec{}
	tests := []struct {
		name     string
		block    blocks.Block
		expected string
		valid    bool
	}{
		{
			"with emoji and color",
			&blocks.CalloutBlock{Callout: blocks.CalloutData{Emoji: "💡", Color: "blue"}},
			`<callout emoji="💡" color="blue">`,
			true,
		},
		{
			"emoji only",
			&blocks.CalloutBlock{Callout: blocks.CalloutData{Emoji: "💡"}},
			`<callout emoji="💡">`,
			true,
		},
		{
			"no attributes",
			&blocks.CalloutBlock{Callout: blocks.CalloutData{}},
			`<callout>`,
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