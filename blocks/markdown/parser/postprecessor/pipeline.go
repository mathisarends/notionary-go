package postprocessor

import (
	blocks "github.com/mathisbot/notionary-go/blocks"
)


type BlockPostProcessor struct {
	processors []PostProcessor
}

func NewBlockPostProcessor() *BlockPostProcessor {
	return &BlockPostProcessor{}
}

func (p *BlockPostProcessor) Register(processor PostProcessor) {
	p.processors = append(p.processors, processor)
}

func (p *BlockPostProcessor) Process(blocks []blocks.RichTextProvider) []blocks.RichTextProvider {
	result := blocks
	for _, processor := range p.processors {
		result = processor.Process(result)
	}
	return result
}