package goopenai

import (
	"net/http"
)

type (
	Client struct {
		client       http.Request
		apiKey       string
		organization string
	}
)

func OpenAI(apiKey string, organization string) Client {
	return Client{
		apiKey:       apiKey,
		organization: organization,
	}
}
