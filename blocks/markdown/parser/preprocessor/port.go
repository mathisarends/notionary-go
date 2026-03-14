package preprocessor

type PreProcessor interface {
	Process(markdownText string) (string, error)
}