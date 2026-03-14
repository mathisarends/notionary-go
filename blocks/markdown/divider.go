package markdown

import (
	"strings"

	"github.com/mathisbot/notionary-go/blocks"
)

type DividerParser struct {
	BaseParser
}

func (p *DividerParser) Parse(line string) (any, bool) {
	trimmed := strings.TrimSpace(line)
	if trimmed == "---" || trimmed == "***" || trimmed == "___" {
		return blocks.DividerBlock{}, true
	}
	return p.Next(line)
}