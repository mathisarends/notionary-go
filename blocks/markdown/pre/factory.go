package pre

func CreateMarkdownToRichTextPreProcessor() *MarkdownPreProcessor {
	p := NewMarkdownPreProcessor()
	p.Register(NewIndentationNormalizer())
	p.Register(NewColumnSyntaxPreProcessor())
	return p
}