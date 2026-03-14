package markdown

import (
	"fmt"
	"strings"

	"github.com/mathisbot/notionary-go/blocks"
	elements "github.com/mathisbot/notionary-go/blocks/markdown/elements"
	post "github.com/mathisbot/notionary-go/blocks/markdown/parser/postprocessing"
	postprocessor "github.com/mathisbot/notionary-go/blocks/markdown/renderer/postprecessor"
)

type Port interface {
	ToBlocks(markdownText string) ([]any, error)
	ToMarkdown(blocks []blocks.Block) string
}

type Converter struct {
	lineParser     elements.Parser
	postProcessor  *post.BlockPostProcessor
	renderPipeline *postprocessor.Pipeline
}

func NewConverter() *Converter {
	return &Converter{
		lineParser:     newLineParserFromRegistry(),
		postProcessor:  post.CreateMarkdownToRichTextPostProcessor(),
		renderPipeline: postprocessor.NewMarkdownRenderingPipeline(),
	}
}

func NewLineParser() elements.Parser {
	return newLineParserFromRegistry()
}

func Parse(markdownText string) ([]any, error) {
	return NewConverter().ToBlocks(markdownText)
}

func Render(blocksToRender []blocks.Block) string {
	return NewConverter().ToMarkdown(blocksToRender)
}

func (c *Converter) ToBlocks(markdownText string) ([]any, error) {
	if strings.TrimSpace(markdownText) == "" {
		return nil, nil
	}

	lines := strings.Split(markdownText, "\n")
	parsed := make([]any, 0, len(lines))
	providers := make([]blocks.RichTextProvider, 0, len(lines))

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		block, ok := c.lineParser.Parse(line)
		if !ok {
			continue
		}
		parsed = append(parsed, block)
		if provider, ok := block.(blocks.RichTextProvider); ok {
			providers = append(providers, provider)
		}
	}

	if len(providers) > 0 {
		c.postProcessor.Process(providers)
	}

	return parsed, nil
}

func (c *Converter) ToMarkdown(blocksToRender []blocks.Block) string {
	rendered := elements.Render(blocksToRender)
	return c.renderPipeline.Process(rendered)
}

func newLineParserFromRegistry() elements.Parser {
	bulleted := mustSimpleSyntax(BulletedList)
	numbered := mustSimpleSyntax(NumberedList)
	return elements.NewLineParserWithPatterns(bulleted.Pattern, numbered.Pattern)
}

func mustSimpleSyntax(key RegistryKey) SimpleSyntax {
	def, ok := Registry[key]
	if !ok {
		panic(fmt.Sprintf("syntax registry entry missing for key: %s", key))
	}
	syntax, ok := def.(SimpleSyntax)
	if !ok {
		panic(fmt.Sprintf("syntax registry entry %s is not SimpleSyntax", key))
	}
	return syntax
}
