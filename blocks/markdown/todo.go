package markdown

import (
	"strings"

	"github.com/mathisbot/notionary-go/blocks"
)

type ToDoParser struct {
	BaseParser
}

func (p *ToDoParser) Parse(line string) (any, bool) {
	if strings.HasPrefix(line, "- [ ] ") {
		return blocks.ToDoBlock{
			ToDo: blocks.ToDoData{
				RichText: toRichText(strings.TrimPrefix(line, "- [ ] ")),
				Checked:  false,
			},
		}, true
	}
	if strings.HasPrefix(line, "- [x] ") || strings.HasPrefix(line, "- [X] ") {
		return blocks.ToDoBlock{
			ToDo: blocks.ToDoData{
				RichText: toRichText(strings.TrimPrefix(line, "- [x] ")),
				Checked:  true,
			},
		}, true
	}
	return p.Next(line)
}