package handlers

import "github.com/mathisbot/notionary-go/blocks"

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
