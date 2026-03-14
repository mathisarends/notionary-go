package richtext

type Grammar struct {
	BoldWrapper           string
	ItalicWrapper         string
	StrikethroughWrapper  string
	UnderlineWrapper      string
	CodeWrapper           string
	InlineEquationWrapper string
	LinkPrefix            string
	LinkMiddle            string
	LinkSuffix            string
	ColorPrefix           string
	ColorMiddle           string
	ColorSuffix           string
}

func DefaultGrammar() Grammar {
	return Grammar{
		BoldWrapper:           "**",
		ItalicWrapper:         "*",
		StrikethroughWrapper:  "~~",
		UnderlineWrapper:      "__",
		CodeWrapper:           "`",
		InlineEquationWrapper: "$",
		LinkPrefix:            "[",
		LinkMiddle:            "](",
		LinkSuffix:            ")",
		ColorPrefix:           "{color:",
		ColorMiddle:           "}",
		ColorSuffix:           "{/color}",
	}
}