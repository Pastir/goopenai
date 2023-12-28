package goopenai

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

type (
	Client struct {
		ClientCreate http.Request
		ApiKey       string
		BearerToken  string
		Organization string
	}
)

func OpenAI(apiKey string, bearerToken string, organization string) Client {
	return Client{
		BearerToken:  bearerToken,
		ApiKey:       apiKey,
		Organization: organization,
	}
}

func (c *Client) Create(ctx context.Context, create OpenAICreate) (OpenAIResponse, error) {
	rJson, err := json.Marshal(create)
	if err != nil {
		return nil, err
	}

	rBody := bytes.NewReader(rJson)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, create.GetMethod(), create.GetEndpoint(), rBody)
	if err != nil {
		return nil, err
	}

	sh := strings.Builder{}
	sh.WriteString("Bearer ")
	sh.WriteString(c.ApiKey)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", sh.String())

	if c.Organization != "" {
		req.Header.Add("OpenAI-Organization", c.Organization)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		var errJson ResponseError
		err = json.Unmarshal(resBody, &errJson)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(errJson.Err.Message)
	}

	response := create.ResponseStruct()

	err = json.Unmarshal(resBody, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
