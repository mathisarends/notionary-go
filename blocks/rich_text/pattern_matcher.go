package richtext

import (
	"regexp"

	"github.com/mathisbot/notionary-go/blocks"
)

type registeredPattern struct {
	handler Handler
	re      *regexp.Regexp
}

type PatternMatcher struct {
	patterns []registeredPattern
}

func NewPatternMatcher() *PatternMatcher {
	return &PatternMatcher{}
}

func (m *PatternMatcher) Register(re *regexp.Regexp, handler Handler) {
	m.patterns = append(m.patterns, registeredPattern{handler, re})
}

type patternMatch struct {
	handler Handler
	match   []string
	start   int
	end     int
}

func (m *PatternMatcher) FindEarliest(text string) *patternMatch {
	var earliest *patternMatch

	for _, p := range m.patterns {
		loc := p.re.FindStringSubmatchIndex(text)
		if loc == nil {
			continue
		}
		start, end := loc[0], loc[1]
		if earliest == nil || start < earliest.start {
			submatches := p.re.FindStringSubmatch(text)
			earliest = &patternMatch{
				handler: p.handler,
				match:   submatches,
				start:   start,
				end:     end,
			}
		}
	}
	return earliest
}

func (m *PatternMatcher) Split(text string) []blocks.RichText {
	var segments []blocks.RichText
	remaining := text

	for remaining != "" {
		match := m.FindEarliest(remaining)
		if match == nil {
			segments = append(segments, blocks.RichTextFromPlainText(remaining))
			break
		}
		if match.start > 0 {
			segments = append(segments, blocks.RichTextFromPlainText(remaining[:match.start]))
		}
		segments = append(segments, match.handler.Handle(match.match))
		remaining = remaining[match.end:]
	}
	return segments
}
