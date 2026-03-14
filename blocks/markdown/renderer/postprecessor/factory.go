package postprocessor

func NewMarkdownRenderingPipeline() *Pipeline {
	p := NewPipeline()
	p.Register(NewNumberedListPlaceholderReplacer())
	return p
}
