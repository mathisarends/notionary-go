package richtext

import (
	"strings"

	"github.com/mathisbot/notionary-go/blocks"
	"github.com/mathisbot/notionary-go/blocks/rich_text/handlers"
)

type Converter struct {
	matcher *PatternMatcher
}

func NewConverter(matcher *PatternMatcher) *Converter {
	return &Converter{matcher: matcher}
}

func (c *Converter) ToRichText(text string) []blocks.RichText {
	if text == "" {
		return nil
	}
	return c.matcher.Split(text)
}

func (c *Converter) ToMarkdown(richTexts []blocks.RichText) string {
	return c.ToMarkdownWithGrammar(richTexts, DefaultGrammar())
}

func (c *Converter) ToMarkdownWithGrammar(richTexts []blocks.RichText, grammar Grammar) string {
	if len(richTexts) == 0 {
		return ""
	}

	var sb strings.Builder
	for _, rt := range richTexts {
		sb.WriteString(richTextToMarkdown(rt, grammar))
	}
	return sb.String()
}

func richTextToMarkdown(rt blocks.RichText, grammar Grammar) string {
	base := rt.PlainText
	if rt.Text != nil {
		base = rt.Text.Content
	}

	if rt.Type == blocks.RichTextTypeEquation {
		return grammar.InlineEquationWrapper + base + grammar.InlineEquationWrapper
	}

	if base == "" {
		return ""
	}

	if rt.Annotations != nil {
		if rt.Annotations.Code {
			base = grammar.CodeWrapper + base + grammar.CodeWrapper
		}
		if rt.Annotations.Bold {
			base = handlers.BoldToMarkdown(base, grammar.BoldWrapper)
		}
		if rt.Annotations.Italic {
			base = grammar.ItalicWrapper + base + grammar.ItalicWrapper
		}
		if rt.Annotations.Strikethrough {
			base = grammar.StrikethroughWrapper + base + grammar.StrikethroughWrapper
		}
		if rt.Annotations.Underline {
			base = grammar.UnderlineWrapper + base + grammar.UnderlineWrapper
		}
	}

	if rt.Text != nil && rt.Text.Link != nil && rt.Text.Link.URL != "" {
		base = handlers.LinkToMarkdown(base, rt.Text.Link.URL, grammar.LinkPrefix, grammar.LinkMiddle, grammar.LinkSuffix)
	}

	if rt.Annotations != nil {
		base = handlers.ColorToMarkdown(base, rt.Annotations.Color, grammar.ColorPrefix, grammar.ColorMiddle, grammar.ColorSuffix)
	}

	return base
}
