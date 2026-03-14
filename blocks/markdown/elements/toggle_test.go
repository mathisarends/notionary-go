package elements

import (
	"testing"

	"github.com/mathisbot/notionary-go/blocks"
)

func TestToggleCodec_Parse(t *testing.T) {
	c := &ToggleCodec{}
	tests := []struct {
		name  string
		line  string
		valid bool
	}{
		{"with title", `<toggle title="Click to expand">`, true},
		{"without title", `<toggle>`, true},
		{"with level should not match toggle", `<toggle title="Section" level="2">`, false},
		{"invalid", `toggle title="foo"`, false},
		{"empty", ``, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, ok := c.Parse(tt.line)
			if ok != tt.valid {
				t.Fatalf("expected ok=%v, got %v", tt.valid, ok)
			}
			if tt.valid {
				if _, ok := block.(*blocks.ToggleBlock); !ok {
					t.Fatal("expected *blocks.ToggleBlock")
				}
			}
		})
	}
}

func TestToggleCodec_Render(t *testing.T) {
	c := &ToggleCodec{}
	tests := []struct {
		name     string
		block    blocks.Block
		expected string
		valid    bool
	}{
		{
			"with title",
			&blocks.ToggleBlock{Toggle: blocks.ToggleData{Title: toRichText("Click to expand")}},
			`<toggle title="Click to expand">`,
			true,
		},
		{
			"without title",
			&blocks.ToggleBlock{Toggle: blocks.ToggleData{}},
			`<toggle>`,
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

func TestToggleableHeadingCodec_Parse(t *testing.T) {
	c := &ToggleableHeadingCodec{}
	tests := []struct {
		name      string
		line      string
		valid     bool
		wantLevel int
	}{
		{"level 1", `<toggle title="Section" level="1">`, true, 1},
		{"level 2", `<toggle title="Section" level="2">`, true, 2},
		{"level 3", `<toggle title="Section" level="3">`, true, 3},
		{"no level", `<toggle title="Section">`, false, 0},
		{"invalid level", `<toggle title="Section" level="4">`, false, 0},
		{"empty", ``, false, 0},
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
			b, ok := block.(*blocks.ToggleableHeadingBlock)
			if !ok {
				t.Fatal("expected *blocks.ToggleableHeadingBlock")
			}
			if b.ToggleableHeading.Level != tt.wantLevel {
				t.Errorf("expected level=%d, got %d", tt.wantLevel, b.ToggleableHeading.Level)
			}
		})
	}
}

func TestToggleableHeadingCodec_Render(t *testing.T) {
	c := &ToggleableHeadingCodec{}
	tests := []struct {
		name     string
		block    blocks.Block
		expected string
		valid    bool
	}{
		{
			"level 2",
			&blocks.ToggleableHeadingBlock{ToggleableHeading: blocks.ToggleableHeadingData{Title: toRichText("Section"), Level: 2}},
			`<toggle title="Section" level="2">`,
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