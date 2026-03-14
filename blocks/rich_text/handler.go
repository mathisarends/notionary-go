package richtext

import "github.com/mathisbot/notionary-go/blocks"

type Handler interface {
	Tag() string // fast prefix check before regex
	Handle(match []string) blocks.RichText
}