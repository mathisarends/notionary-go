package post

import legacy "github.com/mathisbot/notionary-go/blocks/markdown/parser/postprocessing"

type PostProcessor = legacy.PostProcessor
type BlockPostProcessor = legacy.BlockPostProcessor
type RichTextLengthTruncationPostProcessor = legacy.RichTextLengthTruncationPostProcessor

func NewBlockPostProcessor() *BlockPostProcessor {
	return legacy.NewBlockPostProcessor()
}

func NewRichTextLengthTruncationPostProcessor() *RichTextLengthTruncationPostProcessor {
	return legacy.NewRichTextLengthTruncationPostProcessor()
}

func CreateMarkdownToRichTextPostProcessor() *BlockPostProcessor {
	return legacy.CreateMarkdownToRichTextPostProcessor()
}
