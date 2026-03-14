package richtext

import "github.com/mathisbot/notionary-go/blocks"

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