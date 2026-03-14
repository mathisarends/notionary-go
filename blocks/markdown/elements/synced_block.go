package elements

import (
	"github.com/mathisbot/notionary-go/blocks"
	syntax "github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

type SyncedBlockCodec struct{}

func (c *SyncedBlockCodec) Parse(line string) (blocks.Block, bool) {
	syn, ok := syntax.Registry[syntax.SyncedBlock].(syntax.TagSyntax)
	if !ok {
		return nil, false
	}
	m := syn.Pattern.FindStringSubmatch(line)
	if m == nil {
		return nil, false
	}
	return &blocks.SyncedBlock{
		BaseBlock: blocks.BaseBlock{Type: blocks.BlockTypeSyncedBlock},
		SyncedBlock: blocks.SyncedBlockData{
			ID: m[1],
		},
	}, true
}

func (c *SyncedBlockCodec) Render(block blocks.Block) (string, bool) {
	b, ok := block.(*blocks.SyncedBlock)
	if !ok {
		return "", false
	}
	syn, ok := syntax.Registry[syntax.SyncedBlock].(syntax.TagSyntax)
	if !ok {
		return "", false
	}
	result := syn.OpenTag
	if b.SyncedBlock.ID != "" {
		result += ` id="` + b.SyncedBlock.ID + `"`
	}
	result += ">"
	return result, true
}