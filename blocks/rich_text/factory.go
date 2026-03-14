package richtext

import (
	"github.com/mathisbot/notionary-go/blocks/rich_text/handlers"
)

func NewDefaultConverter() *Converter {
	m := NewPatternMatcher()
	m.Register(handlers.BoldPattern, handlers.BoldHandler{})
	m.Register(handlers.ColorPattern, handlers.ColorHandler{})
	m.Register(handlers.LinkPattern, handlers.LinkHandler{})
	return NewConverter(m)
}
