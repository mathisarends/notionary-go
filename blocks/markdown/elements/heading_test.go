package elements

import (
	"testing"

	"github.com/mathisbot/notionary-go/blocks"
)

func TestHeadingCodec_Parse(t *testing.T) {
	c := &HeadingCodec{}
	tests := []struct {
		name      string
		line      string
		valid     bool
		wantLevel int
		wantType  string
	}{
		{"h1", `# Heading 1`, true, 1, "heading_1"},
		{"h2", `## Heading 2`, true, 2, "heading_2"},
		{"h3", `### Heading 3`, true, 3, "heading_3"},
		{"h4 invalid", `#### Heading 4`, false, 0, ""},
		{"no space", `#Heading`, false, 0, ""},
		{"empty", ``, false, 0, ""},
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
			switch tt.wantLevel {
			case 1:
				if _, ok := block.(*blocks.Heading1Block); !ok {
					t.Fatal("expected *blocks.Heading1Block")
				}
			case 2:
				if _, ok := block.(*blocks.Heading2Block); !ok {
					t.Fatal("expected *blocks.Heading2Block")
				}
			case 3:
				if _, ok := block.(*blocks.Heading3Block); !ok {
					t.Fatal("expected *blocks.Heading3Block")
				}
			}
		})
	}
}

func TestHeadingCodec_Render(t *testing.T) {
	c := &HeadingCodec{}
	tests := []struct {
		name     string
		block    blocks.Block
		expected string
		valid    bool
	}{
		{
			"h1",
			&blocks.Heading1Block{Heading1: blocks.HeadingData{RichText: toRichText("Heading 1")}},
			"# Heading 1",
			true,
		},
		{
			"h2",
			&blocks.Heading2Block{Heading2: blocks.HeadingData{RichText: toRichText("Heading 2")}},
			"## Heading 2",
			true,
		},
		{
			"h3",
			&blocks.Heading3Block{Heading3: blocks.HeadingData{RichText: toRichText("Heading 3")}},
			"### Heading 3",
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