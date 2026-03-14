package elements

import (
	"testing"

	"github.com/mathisbot/notionary-go/blocks"
)

func TestTableCodec_Parse(t *testing.T) {
	c := &TableCodec{}
	tests := []struct {
		name  string
		line  string
		valid bool
	}{
		{"simple row", `| Col A | Col B |`, true},
		{"no pipes", `Col A Col B`, false},
		{"empty", ``, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, ok := c.Parse(tt.line)
			if ok != tt.valid {
				t.Fatalf("expected ok=%v, got %v", tt.valid, ok)
			}
			if tt.valid {
				if _, ok := block.(*blocks.TableBlock); !ok {
					t.Fatal("expected *blocks.TableBlock")
				}
			}
		})
	}
}

func TestTableCodec_Render(t *testing.T) {
	c := &TableCodec{}
	tests := []struct {
		name  string
		block blocks.Block
		valid bool
	}{
		{"valid", &blocks.TableBlock{}, true},
		{"wrong type", &blocks.ParagraphBlock{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, ok := c.Render(tt.block)
			if ok != tt.valid {
				t.Fatalf("expected ok=%v, got %v", tt.valid, ok)
			}
		})
	}
}

func TestTableRowCodec_Parse(t *testing.T) {
	c := &TableRowCodec{}
	tests := []struct {
		name  string
		line  string
		valid bool
	}{
		{"separator", `|-------|-------|`, true},
		{"with colons", `|:------|------:|`, true},
		{"data row should not match", `| Col A | Col B |`, false},
		{"empty", ``, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, ok := c.Parse(tt.line)
			if ok != tt.valid {
				t.Fatalf("expected ok=%v, got %v", tt.valid, ok)
			}
			if tt.valid {
				if _, ok := block.(*blocks.TableRowBlock); !ok {
					t.Fatal("expected *blocks.TableRowBlock")
				}
			}
		})
	}
}

func TestTableRowCodec_Render(t *testing.T) {
	c := &TableRowCodec{}
	tests := []struct {
		name  string
		block blocks.Block
		valid bool
	}{
		{"valid", &blocks.TableRowBlock{}, true},
		{"wrong type", &blocks.ParagraphBlock{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, ok := c.Render(tt.block)
			if ok != tt.valid {
				t.Fatalf("expected ok=%v, got %v", tt.valid, ok)
			}
		})
	}
}