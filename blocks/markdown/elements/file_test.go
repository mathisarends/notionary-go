package elements

import (
	"testing"

	"github.com/mathisbot/notionary-go/blocks"
)

func TestFileCodec_Parse(t *testing.T) {
	c := &FileCodec{}
	tests := []struct {
		name  string
		line  string
		valid bool
	}{
		{"with name", `<file src="https://example.com/data.zip" name="data.zip">`, true},
		{"without name", `<file src="https://example.com/data.zip">`, true},
		{"invalid", `file src="https://example.com/data.zip"`, false},
		{"empty", ``, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, ok := c.Parse(tt.line)
			if ok != tt.valid {
				t.Fatalf("expected ok=%v, got %v", tt.valid, ok)
			}
			if tt.valid {
				if _, ok := block.(*blocks.FileBlock); !ok {
					t.Fatal("expected *blocks.FileBlock")
				}
			}
		})
	}
}

func TestFileCodec_Render(t *testing.T) {
	c := &FileCodec{}
	tests := []struct {
		name     string
		block    blocks.Block
		expected string
		valid    bool
	}{
		{
			"with name",
			&blocks.FileBlock{File: blocks.FileData{URL: "https://example.com/data.zip", Name: "data.zip"}},
			`<file src="https://example.com/data.zip" name="data.zip">`,
			true,
		},
		{
			"without name",
			&blocks.FileBlock{File: blocks.FileData{URL: "https://example.com/data.zip"}},
			`<file src="https://example.com/data.zip">`,
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