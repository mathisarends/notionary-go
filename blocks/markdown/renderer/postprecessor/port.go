package postprocessor

type PostProcessor interface {
	Process(markdownText string) string
}
