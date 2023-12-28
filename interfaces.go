package goopenai

type (
	OpenAICreate interface {
		Create() ([]byte, error)
		GetEndpoint() string
		GetMethod() string
		ResponseStruct() OpenAIResponse
	}

	OpenAIResponse interface {
		GetContent() []string
		GetTotalTokens() int
	}
)
