package blocks

type ParagraphData struct {
	RichText []RichText `json:"rich_text"`
	Color    BlockColor `json:"color"`
	Children []Block    `json:"-"` // wird separat geladen
}

type ParagraphBlock struct {
	BaseBlock
	Paragraph ParagraphData `json:"paragraph"`
}

func (b *ParagraphBlock) RichTextRefs() []*[]RichText {
	return []*[]RichText{&b.Paragraph.RichText}
}

type HeadingData struct {
	RichText     []RichText `json:"rich_text"`
	Color        BlockColor `json:"color"`
	IsToggleable bool       `json:"is_toggleable"`
}

type Heading1Block struct {
	BaseBlock
	Heading1 HeadingData `json:"heading_1"`
}

func (b *Heading1Block) RichTextRefs() []*[]RichText {
	return []*[]RichText{&b.Heading1.RichText}
}

type Heading2Block struct {
	BaseBlock
	Heading2 HeadingData `json:"heading_2"`
}

func (b *Heading2Block) RichTextRefs() []*[]RichText {
	return []*[]RichText{&b.Heading2.RichText}
}

type Heading3Block struct {
	BaseBlock
	Heading3 HeadingData `json:"heading_3"`
}

func (b *Heading3Block) RichTextRefs() []*[]RichText {
	return []*[]RichText{&b.Heading3.RichText}
}

type CodeData struct {
	RichText []RichText `json:"rich_text"`
	Caption  []RichText `json:"caption"`
	Language string     `json:"language"`
}

type CodeBlock struct {
	BaseBlock
	Code CodeData `json:"code"`
}

func (b *CodeBlock) RichTextRefs() []*[]RichText {
	return []*[]RichText{&b.Code.RichText, &b.Code.Caption}
}

type ListItemData struct {
	RichText []RichText `json:"rich_text"`
	Color    BlockColor `json:"color"`
}

type BulletedListItemBlock struct {
	BaseBlock
	BulletedListItem ListItemData `json:"bulleted_list_item"`
}

func (b *BulletedListItemBlock) RichTextRefs() []*[]RichText {
	return []*[]RichText{&b.BulletedListItem.RichText}
}

type NumberedListItemBlock struct {
	BaseBlock
	NumberedListItem ListItemData `json:"numbered_list_item"`
}

func (b *NumberedListItemBlock) RichTextRefs() []*[]RichText {
	return []*[]RichText{&b.NumberedListItem.RichText}
}

type QuoteData struct {
	RichText []RichText `json:"rich_text"`
	Color    BlockColor `json:"color"`
}

type QuoteBlock struct {
	BaseBlock
	Quote QuoteData `json:"quote"`
}

func (b *QuoteBlock) RichTextRefs() []*[]RichText {
	return []*[]RichText{&b.Quote.RichText}
}

type CalloutData struct {
	RichText []RichText `json:"rich_text"`
	Color    BlockColor `json:"color"`
}

type CalloutBlock struct {
	BaseBlock
	Callout CalloutData `json:"callout"`
}

func (b *CalloutBlock) RichTextRefs() []*[]RichText {
	return []*[]RichText{&b.Callout.RichText}
}

type ToggleData struct {
	RichText []RichText `json:"rich_text"`
	Color    BlockColor `json:"color"`
}

type ToggleBlock struct {
	BaseBlock
	Toggle ToggleData `json:"toggle"`
}

func (b *ToggleBlock) RichTextRefs() []*[]RichText {
	return []*[]RichText{&b.Toggle.RichText}
}

type ToDoData struct {
	RichText []RichText `json:"rich_text"`
	Checked  bool       `json:"checked"`
	Color    BlockColor `json:"color"`
}

type ToDoBlock struct {
	BaseBlock
	ToDo ToDoData `json:"to_do"`
}

func (b *ToDoBlock) RichTextRefs() []*[]RichText {
	return []*[]RichText{&b.ToDo.RichText}
}

type DividerBlock struct {
	BaseBlock
	Divider struct{} `json:"divider"`
}

type FileData struct {
	Type     string `json:"type"` // "external" | "file"
	URL      string `json:"url,omitempty"`
	External *struct {
		URL string `json:"url"`
	} `json:"external,omitempty"`
	File *struct {
		URL        string `json:"url"`
		ExpiryTime string `json:"expiry_time"`
	} `json:"file,omitempty"`
	Caption []RichText `json:"caption"`
}

type ImageBlock struct {
	BaseBlock
	Image FileData `json:"image"`
}

func (b *ImageBlock) RichTextRefs() []*[]RichText {
	return []*[]RichText{&b.Image.Caption}
}

type VideoBlock struct {
	BaseBlock
	Video FileData `json:"video"`
}

func (b *VideoBlock) RichTextRefs() []*[]RichText {
	return []*[]RichText{&b.Video.Caption}
}

type TableData struct {
	TableWidth      int  `json:"table_width"`
	HasColumnHeader bool `json:"has_column_header"`
	HasRowHeader    bool `json:"has_row_header"`
}

type TableBlock struct {
	BaseBlock
	Table TableData `json:"table"`
}

type TableRowData struct {
	Cells [][]RichText `json:"cells"`
}

type TableRowBlock struct {
	BaseBlock
	TableRow TableRowData `json:"table_row"`
}

type ChildPageBlock struct {
	BaseBlock
	ChildPage struct {
		Title string `json:"title"`
	} `json:"child_page"`
}

type ChildDatabaseBlock struct {
	BaseBlock
	ChildDatabase struct {
		Title string `json:"title"`
	} `json:"child_database"`
}

type UnsupportedBlock struct {
	BaseBlock
}