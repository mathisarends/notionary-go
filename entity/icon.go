package entity

type IconType string

const (
	IconTypeEmoji    IconType = "emoji"
	IconTypeExternal IconType = "external"
	IconTypeFile     IconType = "file"
)

type Icon struct {
	Type     IconType `json:"type"`
	Emoji    *string  `json:"emoji,omitempty"`
	External *URL     `json:"external,omitempty"`
	File     *FileRef `json:"file,omitempty"`
}

func (i *Icon) EmojiValue() (string, bool) {
	if i.Type != IconTypeEmoji || i.Emoji == nil {
		return "", false
	}
	return *i.Emoji, true
}

func (i *Icon) ExternalURL() (string, bool) {
	switch i.Type {
	case IconTypeExternal:
		if i.External != nil {
			return i.External.URL, true
		}
	case IconTypeFile:
		if i.File != nil {
			return i.File.URL, true
		}
	}
	return "", false
}