package goopenai

// Docs https://platform.openai.com/docs/models/model-endpoint-compatibility
const (
	mainURL = "https://api.openai.com/v1/"

	//All models except gpt-3.5-turbo-0301 supported. retrieval tool requires gpt-4-1106-preview or gpt-3.5-turbo-1106.
	endpointAssistants = mainURL + "assistants"

	// whisper-1
	endpointAudioTranscriptions = mainURL + "audio/transcriptions"

	// whisper-1
	endpointAudioTranslations = mainURL + "audio/translations"

	// tts-1, tts-1-hd
	endpointAudioSpeech = mainURL + "audio/speech"

	// gpt-4 and dated model releases, gpt-4-1106-preview, gpt-4-vision-preview, gpt-4-32k and dated model releases,
	// gpt-3.5-turbo and dated model releases, gpt-3.5-turbo-16k and dated model releases,
	// fine-tuned versions of gpt-3.5-turbo
	endpointChatCompletions = mainURL + "chat/completions"

	// (Legacy)	gpt-3.5-turbo-instruct, babbage-002, davinci-002
	endpointCompletions = mainURL + "completions"

	// text-embedding-ada-002
	endpointEmbeddings = mainURL + "embeddings"

	// gpt-3.5-turbo, babbage-002, davinci-002
	endpointFineTuningJobs = mainURL + "fine_tuning/jobs"

	// text-moderation-stable, text-moderation-latest
	endpointModerations = mainURL + "moderations"

	// dall-e-2, dall-e-3
	endpointImagesGenerations = mainURL + "images/generations"
)
