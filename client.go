package goopenai

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	Client struct {
		ClientCreate http.Request
		ApiKey       string
		Organization string
	}
)

func OpenAI(apiKey string, organization string) Client {
	return Client{
		ApiKey:       apiKey,
		Organization: organization,
	}
}

func (c *Client) Create(ctx context.Context, create OpenAICreate) string {
	rJson, err := json.Marshal(create)
	if err != nil {
		//return nil, err
	}
	return string(rJson)

}
