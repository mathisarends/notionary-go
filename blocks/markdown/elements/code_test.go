package elements

import (
	"testing"

	"github.com/mathisbot/notionary-go/blocks"
)

func TestCodeCodec_Parse(t *testing.T) {
	c := &CodeCodec{}
	tests := []struct {
		name         string
		line         string
		valid        bool
		wantLanguage string
	}{
		{"with language", "```python", true, "python"},
		{"without language", "```", true, ""},
		{"with leading spaces", "   ```go", true, "go"},
		{"old tag syntax is invalid", `<code lang="python">`, false, ""},
		{"invalid", "code", false, ""},
		{"empty", "", false, ""},
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

			b, ok := block.(*blocks.CodeBlock)
			if !ok {
				t.Fatal("expected *blocks.CodeBlock")
			}
			if b.Code.Language != tt.wantLanguage {
				t.Errorf("expected language %q, got %q", tt.wantLanguage, b.Code.Language)
			}
		})
	}
}

func TestCodeCodec_Render(t *testing.T) {
	c := &CodeCodec{}
	tests := []struct {
		name     string
		block    blocks.Block
		expected string
		valid    bool
	}{
		{
			"with language",
			&blocks.CodeBlock{Code: blocks.CodeData{Language: "python"}},
			"```python",
			true,
		},
		{
			"without language",
			&blocks.CodeBlock{Code: blocks.CodeData{}},
			"```",
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
