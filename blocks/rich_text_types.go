package blocks

type RichTextType string

const (
	RichTextTypeText     RichTextType = "text"
	RichTextTypeMention  RichTextType = "mention"
	RichTextTypeEquation RichTextType = "equation"
)

type Annotations struct {
	Bold          bool       `json:"bold"`
	Italic        bool       `json:"italic"`
	Strikethrough bool       `json:"strikethrough"`
	Underline     bool       `json:"underline"`
	Code          bool       `json:"code"`
	Color         BlockColor `json:"color"`
}

type TextContent struct {
	Content string `json:"content"`
	Link    *struct {
		URL string `json:"url"`
	} `json:"link,omitempty"`
}

type RichText struct {
	Type        RichTextType `json:"type"`
	Text        *TextContent `json:"text,omitempty"`
	Annotations *Annotations `json:"annotations,omitempty"`
	PlainText   string       `json:"plain_text"`
	Href        *string      `json:"href,omitempty"`
}

func RichTextFromPlainText(text string) RichText {
	return RichText{
		Type:      RichTextTypeText,
		Text:      &TextContent{Content: text},
		PlainText: text,
	}
}
