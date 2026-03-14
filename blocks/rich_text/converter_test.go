package richtext

import (
	"testing"

	"github.com/mathisbot/notionary-go/blocks"
)

func TestConverter_ToRichText_ParsesMixedSegments(t *testing.T) {
	c := NewDefaultConverter()

	got := c.ToRichText("A **bold** and [site](https://example.com)")

	if len(got) != 4 {
		t.Fatalf("expected 4 segments, got %d", len(got))
	}

	if got[0].PlainText != "A " {
		t.Fatalf("expected first segment to be plain text 'A ', got %q", got[0].PlainText)
	}

	if got[1].PlainText != "bold" || got[1].Annotations == nil || !got[1].Annotations.Bold {
		t.Fatalf("expected second segment to be bold rich text, got %+v", got[1])
	}

	if got[2].PlainText != " and " {
		t.Fatalf("expected third segment to be plain text ' and ', got %q", got[2].PlainText)
	}

	if got[3].Text == nil || got[3].Text.Link == nil || got[3].Text.Link.URL != "https://example.com" {
		t.Fatalf("expected fourth segment to contain link URL, got %+v", got[3])
	}
}

func TestConverter_ToRichText_EmptyReturnsNil(t *testing.T) {
	c := NewDefaultConverter()

	got := c.ToRichText("")
	if got != nil {
		t.Fatalf("expected nil for empty input, got %#v", got)
	}
}

func TestConverter_ToMarkdownWithGrammar_AppliesFormattingOrder(t *testing.T) {
	c := NewDefaultConverter()
	grammar := Grammar{
		BoldWrapper:           "__",
		ItalicWrapper:         "*",
		StrikethroughWrapper:  "~~",
		UnderlineWrapper:      "++",
		CodeWrapper:           "`",
		InlineEquationWrapper: "$",
		LinkPrefix:            "<",
		LinkMiddle:            "|",
		LinkSuffix:            ">",
		ColorPrefix:           "{c:",
		ColorMiddle:           "}",
		ColorSuffix:           "{/c}",
	}

	richTexts := []blocks.RichText{
		{
			Type:      blocks.RichTextTypeText,
			PlainText: "hello",
			Text: &blocks.TextContent{
				Content: "hello",
				Link: &struct {
					URL string `json:"url"`
				}{URL: "https://example.com"},
			},
			Annotations: &blocks.Annotations{
				Code:  true,
				Bold:  true,
				Color: blocks.BlockColorRed,
			},
		},
	}

	got := c.ToMarkdownWithGrammar(richTexts, grammar)
	want := "{c:red}<__`hello`__|https://example.com>{/c}"
	if got != want {
		t.Fatalf("unexpected markdown. want %q, got %q", want, got)
	}
}

func TestConverter_ToMarkdown_EquationUsesInlineWrapper(t *testing.T) {
	c := NewDefaultConverter()
	richTexts := []blocks.RichText{
		{Type: blocks.RichTextTypeEquation, PlainText: "x+y"},
	}

	got := c.ToMarkdown(richTexts)
	if got != "$x+y$" {
		t.Fatalf("expected equation markdown $x+y$, got %q", got)
	}
}
