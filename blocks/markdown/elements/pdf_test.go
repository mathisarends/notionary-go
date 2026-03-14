package elements

import (
	"testing"

	"github.com/mathisbot/notionary-go/blocks"
)

func TestPDFCodec_Parse(t *testing.T) {
	c := &PDFCodec{}
	tests := []struct {
		name  string
		line  string
		valid bool
	}{
		{"with caption", `<pdf src="https://example.com/doc.pdf" caption="My Doc">`, true},
		{"without caption", `<pdf src="https://example.com/doc.pdf">`, true},
		{"invalid", `pdf src="https://example.com/doc.pdf"`, false},
		{"empty", ``, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, ok := c.Parse(tt.line)
			if ok != tt.valid {
				t.Fatalf("expected ok=%v, got %v", tt.valid, ok)
			}
			if tt.valid {
				if _, ok := block.(*blocks.PDFBlock); !ok {
					t.Fatal("expected *blocks.PDFBlock")
				}
			}
		})
	}
}

func TestPDFCodec_Render(t *testing.T) {
	c := &PDFCodec{}
	tests := []struct {
		name     string
		block    blocks.Block
		expected string
		valid    bool
	}{
		{
			"with caption",
			&blocks.PDFBlock{PDF: blocks.FileData{URL: "https://example.com/doc.pdf", Caption: toRichText("My Doc")}},
			`<pdf src="https://example.com/doc.pdf" caption="My Doc">`,
			true,
		},
		{
			"without caption",
			&blocks.PDFBlock{PDF: blocks.FileData{URL: "https://example.com/doc.pdf"}},
			`<pdf src="https://example.com/doc.pdf">`,
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