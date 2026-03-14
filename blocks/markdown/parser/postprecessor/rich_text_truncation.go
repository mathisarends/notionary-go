package postprocessor

import (
	blocks "github.com/mathisbot/notionary-go/blocks"
)

const (
	notionMaxLength = 2000
	ellipsis        = "..."
)

type RichTextLengthTruncationPostProcessor struct {
	maxTextLength int
}

func NewRichTextLengthTruncationPostProcessor() *RichTextLengthTruncationPostProcessor {
	return &RichTextLengthTruncationPostProcessor{maxTextLength: notionMaxLength}
}

func (p *RichTextLengthTruncationPostProcessor) Process(input []blocks.RichTextProvider) []blocks.RichTextProvider {
	for _, block := range input {
		for _, richTextSlice := range block.RichTextRefs() {
			p.truncateRichTextList(richTextSlice)
		}
	}
	return input
}

func (p *RichTextLengthTruncationPostProcessor) truncateRichTextList(richTexts *[]blocks.RichText) {
	for i := range *richTexts {
		rt := &(*richTexts)[i]
		if p.shouldTruncate(rt) {
			p.truncate(rt)
		}
	}
}

func (p *RichTextLengthTruncationPostProcessor) shouldTruncate(rt *blocks.RichText) bool {
	return rt.Type == blocks.RichTextTypeText &&
		rt.Text != nil &&
		len(rt.Text.Content) > p.maxTextLength
}

func (p *RichTextLengthTruncationPostProcessor) truncate(rt *blocks.RichText) {
	cutoff := p.maxTextLength - len(ellipsis)
	rt.Text.Content = rt.Text.Content[:cutoff] + ellipsis
}