package postprocessor

import (
	rendererhandler "github.com/mathisbot/notionary-go/blocks/markdown/renderer/postprecessor/handler"
)

func NewMarkdownRenderingPipeline() *Pipeline {
	p := NewPipeline()
	p.Register(rendererhandler.NewNumberedListPlaceholderReplacer())
	return p
}