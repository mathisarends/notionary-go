package parser

import (
	"strings"

	"github.com/mathisbot/notionary-go/blocks"
	"github.com/mathisbot/notionary-go/blocks/markdown/elements"
	"github.com/mathisbot/notionary-go/blocks/markdown/parser/post"
	"github.com/mathisbot/notionary-go/blocks/markdown/parser/pre"
)

type Parser struct {
	chain *chain
	pre   *pre.MarkdownPreProcessor
	post  *post.BlockPostProcessor
}

func New(
	pre *pre.MarkdownPreProcessor,
	post *post.BlockPostProcessor,
	parsers ...LineParser,
) *Parser {
	return &Parser{
		chain: newChain(parsers...),
		pre:   pre,
		post:  post,
	}
}

func NewDefault() *Parser {
	return New(
		pre.CreateMarkdownToRichTextPreProcessor(),
		post.CreateMarkdownToRichTextPostProcessor(),
		&elements.AudioCodec{},
		&elements.BookmarkCodec{},
		&elements.BreadcrumbCodec{},
		&elements.BulletedListCodec{},
		&elements.CalloutCodec{},
		&elements.CodeCodec{},
		&elements.ColumnCodec{},
		&elements.DividerCodec{},
		&elements.EmbedCodec{},
		&elements.EquationCodec{},
		&elements.FileCodec{},
		&elements.HeadingCodec{},
		&elements.ImageCodec{},
		&elements.NumberedListCodec{},
		&elements.ParagraphCodec{},
		&elements.PDFCodec{},
		&elements.QuoteCodec{},
		&elements.SyncedBlockCodec{},
		&elements.TableCodec{},
		&elements.TableOfContentsCodec{},
		&elements.ToDoCodec{},
		&elements.ToDoDoneCodec{},
		&elements.ToggleCodec{},
		&elements.ToggleableHeadingCodec{},
		&elements.VideoCodec{},
	)
}

func (p *Parser) Parse(markdown string) ([]blocks.Block, error) {
	if markdown == "" {
		return nil, nil
	}

	if p.pre != nil {
		var err error
		markdown, err = p.pre.Process(markdown)
		if err != nil {
			return nil, err
		}
	}

	result, err := p.processLines(markdown)
	if err != nil {
		return nil, err
	}

	if p.post != nil {
		providers := make([]blocks.RichTextProvider, 0, len(result))
		for _, block := range result {
			if provider, ok := block.(blocks.RichTextProvider); ok {
				providers = append(providers, provider)
			}
		}
		if len(providers) > 0 {
			p.post.Process(providers)
		}
	}

	return result, nil
}

func (p *Parser) processLines(text string) ([]blocks.Block, error) {
	lines := strings.Split(text, "\n")
	resultBlocks := make([]blocks.Block, 0, len(lines))

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		if block, ok := p.chain.Parse(line); ok {
			resultBlocks = append(resultBlocks, block)
		}
	}

	return resultBlocks, nil
}