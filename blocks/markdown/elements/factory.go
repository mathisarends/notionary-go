package markdown

import "regexp"

func NewLineParser() Parser {
	return NewLineParserWithPatterns(nil, nil)
}

func NewLineParserWithPatterns(bulletedPattern, numberedPattern *regexp.Regexp) Parser {
	root := NewBulletedListParser(bulletedPattern)
	root.SetNext(NewNumberedListParser(numberedPattern))
	return root
}
