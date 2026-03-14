package elements

import (
	"testing"

	"github.com/mathisbot/notionary-go/blocks"
)

func TestColumnListCodec_Parse(t *testing.T) {
	c := &ColumnListCodec{}
	tests := []struct {
		name  string
		line  string
		valid bool
	}{
		{"valid", `<columns>`, true},
		{"invalid", `columns`, false},
		{"empty", ``, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, ok := c.Parse(tt.line)
			if ok != tt.valid {
				t.Fatalf("expected ok=%v, got %v", tt.valid, ok)
			}
			if tt.valid {
				if _, ok := block.(*blocks.ColumnListBlock); !ok {
					t.Fatal("expected *blocks.ColumnListBlock")
				}
			}
		})
	}
}

func TestColumnListCodec_Render(t *testing.T) {
	c := &ColumnListCodec{}
	tests := []struct {
		name     string
		block    blocks.Block
		expected string
		valid    bool
	}{
		{"valid", &blocks.ColumnListBlock{}, "<columns>", true},
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

func TestColumnCodec_Parse(t *testing.T) {
	c := &ColumnCodec{}
	tests := []struct {
		name      string
		line      string
		valid     bool
		wantRatio string
	}{
		{"with ratio", `<column ratio="0.5">`, true, "0.5"},
		{"without ratio", `<column>`, true, ""},
		{"invalid", `column ratio="0.5"`, false, ""},
		{"empty", ``, false, ""},
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
			b, ok := block.(*blocks.ColumnBlock)
			if !ok {
				t.Fatal("expected *blocks.ColumnBlock")
			}
			if b.Column.Ratio != tt.wantRatio {
				t.Errorf("expected ratio=%q, got %q", tt.wantRatio, b.Column.Ratio)
			}
		})
	}
}

func TestColumnCodec_Render(t *testing.T) {
	c := &ColumnCodec{}
	tests := []struct {
		name     string
		block    blocks.Block
		expected string
		valid    bool
	}{
		{
			"with ratio",
			&blocks.ColumnBlock{Column: blocks.ColumnData{Ratio: "0.5"}},
			`<column ratio="0.5">`,
			true,
		},
		{
			"without ratio",
			&blocks.ColumnBlock{Column: blocks.ColumnData{}},
			`<column>`,
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