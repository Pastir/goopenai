package goopenai

import (
	"net/http"
	"strings"
)

// ImageCreate https://platform.openai.com/docs/api-reference/images/create
type (
	ImageCreate struct {
		// Required
		// A text description of the desired image(s). The maximum length is 1000 characters for dall-e-2 and 4000 characters for dall-e-3.
		Prompt string `json:"prompt"`

		// Optional
		// Defaults to dall-e-2, The model to use for image generation.
		Model string `json:"model,omitempty"`

		// Optional
		// Defaults to 1
		// The number of images to generate. Must be between 1 and 10. For dall-e-3, only n=1 is supported.
		N int `json:"n,omitempty"`

		// Optional
		// Defaults to standard
		// The quality of the image that will be generated. hd creates images with finer details and greater consistency
		//across the image. This param is only supported for dall-e-3.
		Quality string `json:"quality,omitempty"`

		// Optional
		// Defaults to url
		// The format in which the generated images are returned. Must be one of url or b64_json.
		ResponseFormat string `json:"response_format,omitempty"`

		// Optional
		// Defaults to 1024x1024
		// The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024 for dall-e-2.
		// Must be one of 1024x1024, 1792x1024, or 1024x1792 for dall-e-3 models.
		Size string `json:"size,omitempty"`

		// Optional
		// Defaults to vivid
		// The style of the generated images. Must be one of vivid or natural.
		// Vivid causes the model to lean towards generating hyper-real and dramatic images.
		// Natural causes the model to produce more natural, less hyper-real looking images.
		// This param is only supported for dall-e-3.
		Style string `json:"style,omitempty"`

		// Optional
		// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. Learn more.
		User string `json:"user,omitempty"`
	}

	DataResponse struct {
		// The base64-encoded JSON of the generated image, if response_format is b64_json.
		B64Json string `json:"b64_json,omitempty"`

		// The URL of the generated image, if response_format is url (default).
		Url string `json:"url,omitempty"`

		// The prompt that was used to generate the image, if there was any revision to the prompt.
		RevisedPrompt string `json:"revised_prompt,omitempty"`
	}

	ImageCreateResponse struct {
		Created int `json:"created"`

		// Returns a list of image objects.
		Data []DataResponse `json:"data"`
	}
)

func (i *ImageCreate) ResponseStruct() OpenAIResponse {
	var icr *ImageCreateResponse = &ImageCreateResponse{}
	return icr
}

func (i *ImageCreate) Create() ([]byte, error) {
	var s []byte
	return s, nil
}

func (i *ImageCreate) GetEndpoint() string {
	return endpointImagesGenerations
}

func (i *ImageCreate) GetMethod() string {
	return http.MethodPost
}

func (i *ImageCreateResponse) GetContent() []string {
	var strContent []string

	for _, v := range i.Data {
		if strings.Compare(v.B64Json, "") != 0 {
			strContent = append(strContent, v.B64Json)
		} else if strings.Compare(v.Url, "") != 0 {
			strContent = append(strContent, v.Url)
		}
	}

	return strContent
}

func (i *ImageCreateResponse) GetTotalTokens() int {
	return 0
}
