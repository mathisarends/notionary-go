package postprocessor

import legacy "github.com/mathisbot/notionary-go/blocks/markdown/renderer/postprecessor"

type PostProcessor = legacy.PostProcessor
type Pipeline = legacy.Pipeline
type NumberedListPlaceholderReplacer = legacy.NumberedListPlaceholderReplacer

func NewPipeline() *Pipeline {
	return legacy.NewPipeline()
}

func NewNumberedListPlaceholderReplacer() *NumberedListPlaceholderReplacer {
	return legacy.NewNumberedListPlaceholderReplacer()
}

func NewMarkdownRenderingPipeline() *Pipeline {
	return legacy.NewMarkdownRenderingPipeline()
}
