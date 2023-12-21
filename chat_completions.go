package goopenai

const (
	RoleAssistant = "assistant"
	RoleUser      = "user"
	RoleSystem    = "system"
)

// CreateChatCompletions https://platform.openai.com/docs/api-reference/chat
type (
	// ToolCall The tool calls generated by the model, such as function calls.
	ToolCall struct {
		// Required, The ID of the tool call.
		Id string `json:"id"`

		// Required, The type of the tool. Currently, only function is supported.
		Type string `json:"type"`

		// Required, The function that the model called.
		Function struct {
			// Required, The name of the function to call.
			Name      string `json:"name"`
			Arguments string `json:"arguments"`
		}
	}

	// Message A list of messages comprising the conversation so far.
	Message struct {
		// The contents of the system message.
		Content string `json:"content"`

		// The role of the messages author, in this case system
		Role string `json:"role"`

		// An optional name for the participant.
		// Provides the model information to differentiate between participants of the same role.
		Name string `json:"name,omitempty"`

		// This case assistant
		// The tool calls generated by the model, such as function calls.
		ToolCalls []ToolCall `json:"tool_calls,omitempty"`

		// The role of the messages author, in this case tool
		ToolCallId string `json:"tool_call_id,omitempty"`
	}

	Function struct {
		// Optional
		Description string `json:"description,omitempty"`

		// Required
		Name string `json:"name"`

		// Optional
		Parameters any `json:"parameters,omitempty"`
	}

	Tool struct {
		// Required, The type of the tool. Currently, only function is supported.
		Type     string    `json:"type"`
		Function *Function `json:"function"`
	}

	CreateChatCompletions struct {
		// Required
		Messages []Message `json:"messages"`

		// Required, ID of the model to use. See the model endpoint compatibility table for details on which models work with the Chat API.
		Model string `json:"model"`

		// Optional
		FrequencyPenalty int `json:"frequency_penalty,omitempty"`

		// Optional
		LogitBias map[string]string `json:"logit_bias,omitempty"`

		// Optional
		Logprobs bool `json:"logprobs,omitempty"`

		// Optional
		TopLogprobs int `json:"top_logprobs,omitempty"`

		// Optional
		MaxTokens int `json:"max_tokens,omitempty"`

		// Optional
		N int `json:"n,omitempty"`

		// Optional
		PresencePenalty int `json:"presence_penalty,omitempty"`

		// Optional
		ResponseFormat struct {
			// Optional
			Type string `json:"type,omitempty"`
		} `json:"response_format,omitempty"`

		// Optional
		Seed int `json:"seed,omitempty"`

		// Optional
		Stop []string `json:"stop,omitempty"`

		// Optional
		Stream bool `json:"stream,omitempty"`

		// Optional
		Temperature int `json:"temperature,omitempty"`

		// Optional
		TopP int `json:"topP,omitempty"`

		// Optional
		Tools []Tool `json:"tools"`

		// Optional need implementation
		ToolChoice interface{} `json:"tool_choice,omitempty"`

		// Optional
		User string `json:"user,omitempty"`
	}

	UsageResponse struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	}

	ChoicesResponse struct {
		Index    int `json:"index"`
		Messages struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Logprobs     bool   `json:"logprobs,omitempty"`
		FinishReason string `json:"finish_reason,omitempty"`
	}

	ChatCompletionsResponse struct {
		ID                string            `json:"id,omitempty"`
		Object            string            `json:"object,omitempty"`
		Created           int               `json:"created"`
		Model             string            `json:"model"`
		SystemFingerprint string            `json:"system_fingerprint"`
		Choices           []ChoicesResponse `json:"choices"`
		Usage             UsageResponse     `json:"usage"`
	}
)