package postprocessor

type Pipeline struct {
	processors []PostProcessor
}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}

func (p *Pipeline) Register(processor PostProcessor) {
	p.processors = append(p.processors, processor)
}

func (p *Pipeline) Process(markdownText string) string {
	for _, processor := range p.processors {
		markdownText = processor.Process(markdownText)
	}
	return markdownText
}