package elements

import (
	"testing"

	"github.com/mathisbot/notionary-go/blocks"
)

func TestAudioCodec_Parse(t *testing.T) {
	c := &AudioCodec{}
	tests := []struct {
		name  string
		line  string
		valid bool
	}{
		{"with caption", `<audio src="https://example.com/track.mp3" caption="My Track">`, true},
		{"without caption", `<audio src="https://example.com/track.mp3">`, true},
		{"invalid", `audio src="https://example.com/track.mp3"`, false},
		{"empty", ``, false},
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
			b, ok := block.(*blocks.AudioBlock)
			if !ok {
				t.Fatal("expected *blocks.AudioBlock")
			}
			if b.Audio.URL == "" {
				t.Error("expected non-empty URL")
			}
		})
	}
}

func TestAudioCodec_Render(t *testing.T) {
	c := &AudioCodec{}
	tests := []struct {
		name     string
		block    blocks.Block
		expected string
		valid    bool
	}{
		{
			"with caption",
			&blocks.AudioBlock{Audio: blocks.FileData{URL: "https://example.com/track.mp3", Caption: toRichText("My Track")}},
			`<audio src="https://example.com/track.mp3" caption="My Track">`,
			true,
		},
		{
			"without caption",
			&blocks.AudioBlock{Audio: blocks.FileData{URL: "https://example.com/track.mp3"}},
			`<audio src="https://example.com/track.mp3">`,
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