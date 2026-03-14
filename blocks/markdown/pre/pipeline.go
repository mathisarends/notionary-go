package pre

type MarkdownPreProcessor struct {
	processors []PreProcessor
}

func NewMarkdownPreProcessor() *MarkdownPreProcessor {
	return &MarkdownPreProcessor{}
}

func (m *MarkdownPreProcessor) Register(p PreProcessor) {
	m.processors = append(m.processors, p)
}

func (m *MarkdownPreProcessor) Process(markdownText string) (string, error) {
	result := markdownText
	for _, p := range m.processors {
		var err error
		result, err = p.Process(result)
		if err != nil {
			return "", err
		}
	}
	return result, nil
}