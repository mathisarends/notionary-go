package elements

import (
	"testing"

	"github.com/mathisbot/notionary-go/blocks"
)

func TestVideoCodec_Parse(t *testing.T) {
	c := &VideoCodec{}
	tests := []struct {
		name  string
		line  string
		valid bool
	}{
		{"with caption", `<video src="https://example.com/clip.mp4" caption="My Clip">`, true},
		{"without caption", `<video src="https://example.com/clip.mp4">`, true},
		{"invalid", `video src="https://example.com/clip.mp4"`, false},
		{"empty", ``, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, ok := c.Parse(tt.line)
			if ok != tt.valid {
				t.Fatalf("expected ok=%v, got %v", tt.valid, ok)
			}
			if tt.valid {
				if _, ok := block.(*blocks.VideoBlock); !ok {
					t.Fatal("expected *blocks.VideoBlock")
				}
			}
		})
	}
}

func TestVideoCodec_Render(t *testing.T) {
	c := &VideoCodec{}
	tests := []struct {
		name     string
		block    blocks.Block
		expected string
		valid    bool
	}{
		{
			"with caption",
			&blocks.VideoBlock{Video: blocks.FileData{URL: "https://example.com/clip.mp4", Caption: toRichText("My Clip")}},
			`<video src="https://example.com/clip.mp4" caption="My Clip">`,
			true,
		},
		{
			"without caption",
			&blocks.VideoBlock{Video: blocks.FileData{URL: "https://example.com/clip.mp4"}},
			`<video src="https://example.com/clip.mp4">`,
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