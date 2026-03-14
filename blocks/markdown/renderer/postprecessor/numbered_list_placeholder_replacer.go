package postprocessor

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	placeholder    = "__NUM__"
	spacesPerLevel = 4
)

var (
	placeholderPattern  = regexp.MustCompile(`^\s*` + regexp.QuoteMeta(placeholder) + `\.`)
	indentationPattern  = regexp.MustCompile(`^(\s*)` + regexp.QuoteMeta(placeholder) + `\.`)
	contentPattern      = regexp.MustCompile(`^\s*` + regexp.QuoteMeta(placeholder) + `\.\s*(.*)`)
	numberedItemPattern = regexp.MustCompile(`^\s*(\d+|[a-z]+|[ivxlcdm]+)\.\s+`)
)

type NumberedListPlaceholderReplacer struct{}

func NewNumberedListPlaceholderReplacer() *NumberedListPlaceholderReplacer {
	return &NumberedListPlaceholderReplacer{}
}

func (r *NumberedListPlaceholderReplacer) Process(markdownText string) string {
	lines := strings.Split(markdownText, "\n")
	result := make([]string, 0, len(lines))
	state := newListNumberingState()

	for i, line := range lines {
		switch {
		case placeholderPattern.MatchString(line):
			result = append(result, r.convertToNumberedItem(line, state))
		case r.isBlankBetweenListItems(lines, i, result):
			continue
		default:
			state.reset()
			result = append(result, line)
		}
	}

	return strings.Join(result, "\n")
}

func (r *NumberedListPlaceholderReplacer) convertToNumberedItem(line string, state *listNumberingState) string {
	indentation := r.extractIndentation(line)
	content := r.extractContent(line)
	nestingLevel := len(indentation) / spacesPerLevel

	state.advanceToLevel(nestingLevel)
	return fmt.Sprintf("%s%s. %s", indentation, state.numberForCurrentLevel(), content)
}

func (r *NumberedListPlaceholderReplacer) extractIndentation(line string) string {
	if m := indentationPattern.FindStringSubmatch(line); m != nil {
		return m[1]
	}
	return ""
}

func (r *NumberedListPlaceholderReplacer) extractContent(line string) string {
	if m := contentPattern.FindStringSubmatch(line); m != nil {
		return m[1]
	}
	return ""
}

func (r *NumberedListPlaceholderReplacer) isBlankBetweenListItems(lines []string, index int, processed []string) bool {
	if strings.TrimSpace(lines[index]) != "" {
		return false
	}
	if len(processed) == 0 || !numberedItemPattern.MatchString(processed[len(processed)-1]) {
		return false
	}
	return index+1 < len(lines) && placeholderPattern.MatchString(lines[index+1])
}