package handlers

import (
	"regexp"
	"strings"

	"github.com/mathisbot/notionary-go/blocks"
)

var ColorPattern = regexp.MustCompile(`\{color:([a-z_]+)\}(.+?)\{/color\}`)

type ColorHandler struct{}

func (ColorHandler) Tag() string { return "{color:" }

func (ColorHandler) Handle(match []string) blocks.RichText {
	content := match[2]
	color := parseColor(match[1])
	return blocks.RichText{
		Type:      blocks.RichTextTypeText,
		PlainText: content,
		Text:      &blocks.TextContent{Content: content},
		Annotations: &blocks.Annotations{
			Color: color,
		},
	}
}

func ColorToMarkdown(text string, color blocks.BlockColor, prefix, middle, suffix string) string {
	if text == "" || color == "" || color == blocks.BlockColorDefault {
		return text
	}
	if prefix == "" {
		prefix = "{color:"
	}
	if middle == "" {
		middle = "}"
	}
	if suffix == "" {
		suffix = "{/color}"
	}
	return prefix + string(color) + middle + text + suffix
}

type ColorGroup struct {
	Color   blocks.BlockColor
	Entries []blocks.RichText
}

func ChunkByColor(richTexts []blocks.RichText) []ColorGroup {
	if len(richTexts) == 0 {
		return nil
	}

	var groups []ColorGroup
	currentColor := extractColor(richTexts[0])
	current := []blocks.RichText{}

	for _, rt := range richTexts {
		c := extractColor(rt)
		if c == currentColor {
			current = append(current, rt)
		} else {
			groups = append(groups, ColorGroup{Color: currentColor, Entries: current})
			currentColor = c
			current = []blocks.RichText{rt}
		}
	}
	return append(groups, ColorGroup{Color: currentColor, Entries: current})
}

func extractColor(rt blocks.RichText) blocks.BlockColor {
	if rt.Annotations != nil {
		return rt.Annotations.Color
	}
	return blocks.BlockColorDefault
}

func parseColor(raw string) blocks.BlockColor {
	candidate := blocks.BlockColor(strings.ToLower(strings.TrimSpace(raw)))
	if isSupportedColor(candidate) {
		return candidate
	}
	return blocks.BlockColorDefault
}

func isSupportedColor(c blocks.BlockColor) bool {
	switch c {
	case blocks.BlockColorDefault,
		blocks.BlockColorBlue,
		blocks.BlockColorBrown,
		blocks.BlockColorGray,
		blocks.BlockColorGreen,
		blocks.BlockColorOrange,
		blocks.BlockColorPink,
		blocks.BlockColorPurple,
		blocks.BlockColorRed,
		blocks.BlockColorYellow,
		blocks.BlockColorBlueBackground,
		blocks.BlockColorBrownBackground,
		blocks.BlockColorGrayBackground,
		blocks.BlockColorGreenBackground,
		blocks.BlockColorOrangeBackground,
		blocks.BlockColorPinkBackground,
		blocks.BlockColorPurpleBackground,
		blocks.BlockColorRedBackground,
		blocks.BlockColorYellowBackground:
		return true
	default:
		return false
	}
}
