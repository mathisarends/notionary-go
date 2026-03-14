package elements

import (
	"github.com/mathisbot/notionary-go/blocks"
)

type BlockCodec interface {
	Render(block blocks.Block) (string, bool)
	Parse(line string) (blocks.Block, bool)
}