package blocks

type RichTextProvider interface {
	RichTextRefs() []*[]RichText
}