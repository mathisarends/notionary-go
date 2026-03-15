package elements

import (
	"testing"

	"github.com/mathisbot/notionary-go/blocks"
)

func TestImageCodec_Parse(t *testing.T) {
	c := &ImageCodec{}
	tests := []struct {
		name  string
		line  string
		valid bool
	}{
		{"with caption", `<image src="https://example.com/photo.png" caption="A photo">`, true},
		{"without caption", `<image src="https://example.com/photo.png">`, true},
		{"invalid", `image src="https://example.com/photo.png"`, false},
		{"empty", ``, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, ok := c.Parse(tt.line)
			if ok != tt.valid {
				t.Fatalf("expected ok=%v, got %v", tt.valid, ok)
			}
			if tt.valid {
				if _, ok := block.(*blocks.ImageBlock); !ok {
					t.Fatal("expected *blocks.ImageBlock")
				}
			}
		})
	}
}

func TestImageCodec_Render(t *testing.T) {
	c := &ImageCodec{}
	tests := []struct {
		name     string
		block    blocks.Block
		expected string
		valid    bool
	}{
		{
			"with caption",
			&blocks.ImageBlock{Image: blocks.FileData{URL: "https://example.com/photo.png", Caption: toRichText("A photo")}},
			`<image src="https://example.com/photo.png" caption="A photo">`,
			true,
		},
		{
			"without caption",
			&blocks.ImageBlock{Image: blocks.FileData{URL: "https://example.com/photo.png"}},
			`<image src="https://example.com/photo.png">`,
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