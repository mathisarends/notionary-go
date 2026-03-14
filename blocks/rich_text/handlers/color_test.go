package handlers

import (
	"testing"

	"github.com/mathisbot/notionary-go/blocks"
)

func TestColorToMarkdown(t *testing.T) {
	t.Run("ignores empty and default color", func(t *testing.T) {
		if got := ColorToMarkdown("x", "", "", "", ""); got != "x" {
			t.Fatalf("expected unchanged text for empty color, got %q", got)
		}
		if got := ColorToMarkdown("x", blocks.BlockColorDefault, "", "", ""); got != "x" {
			t.Fatalf("expected unchanged text for default color, got %q", got)
		}
	})

	t.Run("uses defaults", func(t *testing.T) {
		if got := ColorToMarkdown("x", blocks.BlockColorBlue, "", "", ""); got != "{color:blue}x{/color}" {
			t.Fatalf("expected default color markdown, got %q", got)
		}
	})

	t.Run("uses custom wrappers", func(t *testing.T) {
		if got := ColorToMarkdown("x", blocks.BlockColorBlue, "<c:", ">", "</c>"); got != "<c:blue>x</c>" {
			t.Fatalf("expected custom wrapped color markdown, got %q", got)
		}
	})
}

func TestColorHandler_Handle_ParsesAndNormalizesColor(t *testing.T) {
	match := ColorPattern.FindStringSubmatch("{color:red_background}hello{/color}")
	if len(match) < 3 {
		t.Fatalf("expected regex submatch for color syntax, got %#v", match)
	}

	rt := (ColorHandler{}).Handle(match)

	if rt.PlainText != "hello" {
		t.Fatalf("expected plain text 'hello', got %q", rt.PlainText)
	}
	if rt.Annotations == nil || rt.Annotations.Color != blocks.BlockColorRedBackground {
		t.Fatalf("expected red background color annotation, got %+v", rt.Annotations)
	}
}

func TestChunkByColor_MissingAnnotationColorTreatedAsDefault(t *testing.T) {
	richTexts := []blocks.RichText{
		{Type: blocks.RichTextTypeText, PlainText: "a"},
		{
			Type:        blocks.RichTextTypeText,
			PlainText:   "b",
			Annotations: &blocks.Annotations{Bold: true},
		},
	}

	groups := ChunkByColor(richTexts)

	if len(groups) != 1 {
		t.Fatalf("expected one default-color group for entries without explicit color, got %d groups: %#v", len(groups), groups)
	}
	if groups[0].Color != blocks.BlockColorDefault {
		t.Fatalf("expected group color default, got %q", groups[0].Color)
	}
	if len(groups[0].Entries) != 2 {
		t.Fatalf("expected 2 entries in the single group, got %d", len(groups[0].Entries))
	}
}
