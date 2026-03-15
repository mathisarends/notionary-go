package renderer

import "github.com/mathisbot/notionary-go/blocks"

type ElementRenderer interface {
	Render(block blocks.Block) (string, bool)
}

type chain struct {
	renderers []ElementRenderer
}

func newChain(renderers ...ElementRenderer) *chain {
	return &chain{renderers: renderers}
}

func (c *chain) handle(ctx *Context) {
	for _, r := range c.renderers {
		if result, ok := r.Render(ctx.Block); ok {
			ctx.Result = result
			return
		}
	}
}