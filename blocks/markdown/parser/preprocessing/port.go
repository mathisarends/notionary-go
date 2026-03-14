package pre

type PreProcessor interface {
	Process(markdownText string) (string, error)
}