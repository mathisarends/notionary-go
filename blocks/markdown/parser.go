package markdown

type Parser interface {
	Parse(line string) (any, bool)
	SetNext(p Parser) Parser
}

type BaseParser struct {
	next Parser
}

func (b *BaseParser) SetNext(p Parser) Parser {
	b.next = p
	return p
}

func (b *BaseParser) Next(line string) (any, bool) {
	if b.next != nil {
		return b.next.Parse(line)
	}
	return nil, false
}