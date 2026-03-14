package handlers

import "testing"

func TestBoldToMarkdown(t *testing.T) {
	t.Run("empty text", func(t *testing.T) {
		if got := BoldToMarkdown("", "**"); got != "" {
			t.Fatalf("expected empty string, got %q", got)
		}
	})

	t.Run("uses default wrapper", func(t *testing.T) {
		if got := BoldToMarkdown("abc", ""); got != "**abc**" {
			t.Fatalf("expected default wrapper result, got %q", got)
		}
	})

	t.Run("uses custom wrapper", func(t *testing.T) {
		if got := BoldToMarkdown("abc", "__"); got != "__abc__" {
			t.Fatalf("expected custom wrapper result, got %q", got)
		}
	})
}

func TestBoldHandler_Handle(t *testing.T) {
	match := BoldPattern.FindStringSubmatch("**strong**")
	if len(match) < 2 {
		t.Fatalf("expected regex submatch for bold syntax, got %#v", match)
	}

	rt := (BoldHandler{}).Handle(match)

	if rt.PlainText != "strong" {
		t.Fatalf("expected plain text 'strong', got %q", rt.PlainText)
	}
	if rt.Text == nil || rt.Text.Content != "strong" {
		t.Fatalf("expected text content 'strong', got %+v", rt.Text)
	}
	if rt.Annotations == nil || !rt.Annotations.Bold {
		t.Fatalf("expected bold annotation, got %+v", rt.Annotations)
	}
}
