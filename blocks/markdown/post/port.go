package post

import (
	blocks "github.com/mathisbot/notionary-go/blocks"
)

type PostProcessor interface {
	Process(blocks []blocks.RichTextProvider) []blocks.RichTextProvider
}