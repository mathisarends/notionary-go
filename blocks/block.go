package blocks

import (
	"encoding/json"
	"fmt"
	"time"
)

type Block interface {
	GetID() string
	GetType() BlockType
	GetHasChildren() bool
}

type BaseBlock struct {
	ID             string    `json:"id"`
	Type           BlockType `json:"type"`
	CreatedTime    time.Time `json:"created_time"`
	LastEditedTime time.Time `json:"last_edited_time"`
	Archived       bool      `json:"archived"`
	InTrash        bool      `json:"in_trash"`
	HasChildren    bool      `json:"has_children"`
}

func (b BaseBlock) GetID() string          { return b.ID }
func (b BaseBlock) GetType() BlockType     { return b.Type }
func (b BaseBlock) GetHasChildren() bool   { return b.HasChildren }

// UnmarshalBlock liest type-field und gibt den richtigen Block zurück
func UnmarshalBlock(data []byte) (Block, error) {
	var raw struct {
		Type BlockType `json:"type"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	var block Block
	switch raw.Type {
	case BlockTypeParagraph:
		block = &ParagraphBlock{}
	case BlockTypeHeading1:
		block = &Heading1Block{}
	case BlockTypeHeading2:
		block = &Heading2Block{}
	case BlockTypeHeading3:
		block = &Heading3Block{}
	case BlockTypeBulletedListItem:
		block = &BulletedListItemBlock{}
	case BlockTypeNumberedListItem:
		block = &NumberedListItemBlock{}
	case BlockTypeCode:
		block = &CodeBlock{}
	case BlockTypeQuote:
		block = &QuoteBlock{}
	case BlockTypeCallout:
		block = &CalloutBlock{}
	case BlockTypeToDo:
		block = &ToDoBlock{}
	case BlockTypeToggle:
		block = &ToggleBlock{}
	case BlockTypeDivider:
		block = &DividerBlock{}
	case BlockTypeImage:
		block = &ImageBlock{}
	case BlockTypeVideo:
		block = &VideoBlock{}
	case BlockTypeTable:
		block = &TableBlock{}
	case BlockTypeTableRow:
		block = &TableRowBlock{}
	case BlockTypeChildPage:
		block = &ChildPageBlock{}
	case BlockTypeChildDatabase:
		block = &ChildDatabaseBlock{}
	default:
		block = &UnsupportedBlock{}
	}

	if err := json.Unmarshal(data, block); err != nil {
		return nil, fmt.Errorf("unmarshal block type %q: %w", raw.Type, err)
	}
	return block, nil
}

// BlockList für API responses mit custom unmarshaling
type BlockList struct {
	Results    []Block `json:"results"`
	HasMore    bool    `json:"has_more"`
	NextCursor *string `json:"next_cursor"`
}

func (bl *BlockList) UnmarshalJSON(data []byte) error {
	var raw struct {
		Results    []json.RawMessage `json:"results"`
		HasMore    bool              `json:"has_more"`
		NextCursor *string           `json:"next_cursor"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	bl.HasMore = raw.HasMore
	bl.NextCursor = raw.NextCursor
	bl.Results = make([]Block, 0, len(raw.Results))

	for _, r := range raw.Results {
		block, err := UnmarshalBlock(r)
		if err != nil {
			return err
		}
		bl.Results = append(bl.Results, block)
	}
	return nil
}