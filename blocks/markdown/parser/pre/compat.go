package pre

import legacy "github.com/mathisbot/notionary-go/blocks/markdown/parser/preprocessing"

type PreProcessor = legacy.PreProcessor
type MarkdownPreProcessor = legacy.MarkdownPreProcessor
type IndentationNormalizer = legacy.IndentationNormalizer
type ColumnSyntaxPreProcessor = legacy.ColumnSyntaxPreProcessor

func NewMarkdownPreProcessor() *MarkdownPreProcessor {
	return legacy.NewMarkdownPreProcessor()
}

func NewIndentationNormalizer() *IndentationNormalizer {
	return legacy.NewIndentationNormalizer()
}

func NewColumnSyntaxPreProcessor() *ColumnSyntaxPreProcessor {
	return legacy.NewColumnSyntaxPreProcessor()
}

func CreateMarkdownToRichTextPreProcessor() *MarkdownPreProcessor {
	return legacy.CreateMarkdownToRichTextPreProcessor()
}
