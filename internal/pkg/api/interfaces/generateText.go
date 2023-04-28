package interfaces

type GPTGenerateText interface {
	GenerateText(prompt string) (string, error)
}
