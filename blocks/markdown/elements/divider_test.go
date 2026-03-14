package elements

import (
	"testing"

	"github.com/mathisbot/notionary-go/blocks"
)

func TestDividerCodec_Parse(t *testing.T) {
	c := &DividerCodec{}
	tests := []struct {
		name  string
		line  string
		valid bool
	}{
		{"three dashes", `---`, true},
		{"five dashes", `-----`, true},
		{"two dashes", `--`, false},
		{"empty", ``, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, ok := c.Parse(tt.line)
			if ok != tt.valid {
				t.Fatalf("expected ok=%v, got %v", tt.valid, ok)
			}
			if tt.valid {
				if _, ok := block.(*blocks.DividerBlock); !ok {
					t.Fatal("expected *blocks.DividerBlock")
				}
			}
		})
	}
}

func TestDividerCodec_Render(t *testing.T) {
	c := &DividerCodec{}
	tests := []struct {
		name     string
		block    blocks.Block
		expected string
		valid    bool
	}{
		{"valid", &blocks.DividerBlock{}, "---", true},
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