package markdown

import (
	"github.com/mathisbot/notionary-go/blocks"
	"github.com/mathisbot/notionary-go/blocks/markdown/parser"
	"github.com/mathisbot/notionary-go/blocks/markdown/renderer"
	postprocessor "github.com/mathisbot/notionary-go/blocks/markdown/renderer/postprocessor"
)

type Port interface {
	ToBlocks(markdownText string) ([]any, error)
	ToMarkdown(blocks []blocks.Block) string
}

type Converter struct {
	parser         *parser.Parser
	renderer       *renderer.Renderer
	renderPipeline *postprocessor.Pipeline
}

func NewConverter() *Converter {
	pipeline := postprocessor.NewMarkdownRenderingPipeline()
	return &Converter{
		parser:         parser.NewDefault(),
		renderer:       renderer.NewDefault(pipeline),
		renderPipeline: pipeline,
	}
}

func Parse(markdownText string) ([]any, error) {
	return NewConverter().ToBlocks(markdownText)
}

func Render(blocksToRender []blocks.Block) string {
	return NewConverter().ToMarkdown(blocksToRender)
}

func (c *Converter) ToBlocks(markdownText string) ([]any, error) {
	if c.parser == nil {
		return nil, nil
	}

	parsedBlocks, err := c.parser.Parse(markdownText)
	if err != nil {
		return nil, err
	}

	parsed := make([]any, 0, len(parsedBlocks))
	for _, block := range parsedBlocks {
		parsed = append(parsed, block)
	}

	return parsed, nil
}

func (c *Converter) ToMarkdown(blocksToRender []blocks.Block) string {
	if c.renderer == nil {
		return ""
	}
	rendered, err := c.renderer.Render(blocksToRender, 0)
	if err != nil {
		return ""
	}
	if c.renderPipeline == nil {
		return rendered
	}
	return c.renderPipeline.Process(rendered)
}
