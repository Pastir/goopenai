package goopenai

type (
	OpenAICreate interface {
		Create() ([]byte, error)
	}
)
