package postprocessor

func CreateMarkdownToRichTextPostProcessor() *BlockPostProcessor {
	p := NewBlockPostProcessor()
	p.Register(NewRichTextLengthTruncationPostProcessor())
	return p
}