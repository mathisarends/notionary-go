package elements

import (
	"testing"

	"github.com/mathisbot/notionary-go/blocks"
)

func TestEquationCodec_Parse(t *testing.T) {
	c := &EquationCodec{}
	tests := []struct {
		name  string
		line  string
		valid bool
	}{
		{"valid", `<equation>`, true},
		{"invalid", `equation`, false},
		{"empty", ``, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, ok := c.Parse(tt.line)
			if ok != tt.valid {
				t.Fatalf("expected ok=%v, got %v", tt.valid, ok)
			}
			if tt.valid {
				if _, ok := block.(*blocks.EquationBlock); !ok {
					t.Fatal("expected *blocks.EquationBlock")
				}
			}
		})
	}
}

func TestEquationCodec_Render(t *testing.T) {
	c := &EquationCodec{}
	tests := []struct {
		name     string
		block    blocks.Block
		expected string
		valid    bool
	}{
		{"valid", &blocks.EquationBlock{}, "<equation>", true},
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