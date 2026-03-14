package handlers

import "testing"

func TestLinkToMarkdown(t *testing.T) {
	t.Run("empty text", func(t *testing.T) {
		if got := LinkToMarkdown("", "https://example.com", "[", "](", ")"); got != "" {
			t.Fatalf("expected empty string, got %q", got)
		}
	})

	t.Run("empty url returns plain text", func(t *testing.T) {
		if got := LinkToMarkdown("label", "", "[", "](", ")"); got != "label" {
			t.Fatalf("expected plain text when url is empty, got %q", got)
		}
	})

	t.Run("uses default wrappers", func(t *testing.T) {
		if got := LinkToMarkdown("label", "https://example.com", "", "", ""); got != "[label](https://example.com)" {
			t.Fatalf("expected default markdown link, got %q", got)
		}
	})

	t.Run("uses custom wrappers", func(t *testing.T) {
		if got := LinkToMarkdown("label", "https://example.com", "<", "|", ">"); got != "<label|https://example.com>" {
			t.Fatalf("expected custom wrapped link, got %q", got)
		}
	})
}

func TestLinkHandler_Handle(t *testing.T) {
	match := LinkPattern.FindStringSubmatch("[docs](https://example.com/docs)")
	if len(match) < 3 {
		t.Fatalf("expected regex submatch for link syntax, got %#v", match)
	}

	rt := (LinkHandler{}).Handle(match)

	if rt.PlainText != "docs" {
		t.Fatalf("expected plain text 'docs', got %q", rt.PlainText)
	}
	if rt.Text == nil || rt.Text.Content != "docs" {
		t.Fatalf("expected text content 'docs', got %+v", rt.Text)
	}
	if rt.Text == nil || rt.Text.Link == nil || rt.Text.Link.URL != "https://example.com/docs" {
		t.Fatalf("expected URL to be extracted, got %+v", rt.Text)
	}
}
