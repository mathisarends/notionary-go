package pre

import (
	"math"
	"strings"
)

const (
	spacesPerNestingLevel   = 4
	codeBlockStartDelimiter = "```"
)

type IndentationNormalizer struct{}

func NewIndentationNormalizer() *IndentationNormalizer {
	return &IndentationNormalizer{}
}

func (n *IndentationNormalizer) Process(markdownText string) (string, error) {
	if markdownText == "" {
		return "", nil
	}

	normalized := n.normalizeToMarkdownIndentation(markdownText)
	return normalized, nil
}

func (n *IndentationNormalizer) normalizeToMarkdownIndentation(text string) string {
	lines := strings.Split(text, "\n")
	processed := make([]string, 0, len(lines))
	insideCodeBlock := false

	for _, line := range lines {
		if n.isCodeFence(line) {
			insideCodeBlock = !insideCodeBlock
			processed = append(processed, line)
		} else if insideCodeBlock {
			processed = append(processed, line)
		} else {
			processed = append(processed, n.normalizeToStandardIndentation(line))
		}
	}

	return strings.Join(processed, "\n")
}

func (n *IndentationNormalizer) isCodeFence(line string) bool {
	return strings.HasPrefix(strings.TrimLeft(line, " \t"), codeBlockStartDelimiter)
}

func (n *IndentationNormalizer) normalizeToStandardIndentation(line string) string {
	if strings.TrimSpace(line) == "" {
		return ""
	}

	level := n.roundToNearestIndentationLevel(line)
	content := strings.TrimLeft(line, " \t")
	return strings.Repeat(" ", level*spacesPerNestingLevel) + content
}

func (n *IndentationNormalizer) roundToNearestIndentationLevel(line string) int {
	leading := len(line) - len(strings.TrimLeft(line, " \t"))
	return int(math.Ceil(float64(leading) / float64(spacesPerNestingLevel)))
}