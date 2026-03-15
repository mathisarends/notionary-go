package renderer

import (
	"strings"

	"github.com/mathisbot/notionary-go/blocks"
	elements "github.com/mathisbot/notionary-go/blocks/markdown/elements"
	"github.com/mathisbot/notionary-go/blocks/markdown/renderer/postprocessor"
)

type Renderer struct {
	chain     *chain
	pipeline  *postprocessor.Pipeline
}

func new(pipeline *postprocessor.Pipeline, renderers ...ElementRenderer) *Renderer {
	return &Renderer{
		chain:    newChain(renderers...),
		pipeline: pipeline,
	}
}

func NewDefault(pipeline *postprocessor.Pipeline) *Renderer {
	return new(pipeline,
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
		&elements.TableRowCodec{},
		&elements.TableOfContentsCodec{},
		&elements.ToDoCodec{},
		&elements.ToDoDoneCodec{},
		&elements.ToggleCodec{},
		&elements.ToggleableHeadingCodec{},
		&elements.VideoCodec{},
	)
}

func (r *Renderer) Render(blks []blocks.Block, indentLevel int) (string, error) {
	if len(blks) == 0 {
		return "", nil
	}

	parts := make([]string, 0, len(blks))

	for _, b := range blks {
		ctx := newContext(b, indentLevel, r.Render)
		r.chain.handle(ctx)
		if ctx.Result != "" {
			parts = append(parts, ctx.Result)
		}
	}

	sep := "\n\n"
	if indentLevel > 0 {
		sep = "\n"
	}
	result := strings.Join(parts, sep)

	if r.pipeline != nil {
		result = r.pipeline.Process(result)
	}

	return result, nil
}