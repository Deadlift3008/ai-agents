package clients

import (
	"bytes"
	"deadlift3008/ai-agents/helpers"
	"encoding/json"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type OpenRouterApiResponse struct {
	ID                string      `json:"id"`
	Object            string      `json:"object"`
	Created           int         `json:"created"`
	Model             string      `json:"model"`
	Provider          string      `json:"provider"`
	SystemFingerprint interface{} `json:"system_fingerprint"`
	ServiceTier       interface{} `json:"service_tier"`
	Choices           []struct {
		Index              int         `json:"index"`
		Logprobs           interface{} `json:"logprobs"`
		FinishReason       string      `json:"finish_reason"`
		NativeFinishReason string      `json:"native_finish_reason"`
		Message            struct {
			Role             string      `json:"role"`
			Content          string      `json:"content"`
			Refusal          interface{} `json:"refusal"`
			Reasoning        string      `json:"reasoning"`
			ReasoningDetails []struct {
				Type   string `json:"type"`
				Text   string `json:"text"`
				Format string `json:"format"`
				Index  int    `json:"index"`
			} `json:"reasoning_details"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		PromptTokens        int     `json:"prompt_tokens"`
		CompletionTokens    int     `json:"completion_tokens"`
		TotalTokens         int     `json:"total_tokens"`
		Cost                float64 `json:"cost"`
		IsByok              bool    `json:"is_byok"`
		PromptTokensDetails struct {
			CachedTokens     int `json:"cached_tokens"`
			CacheWriteTokens int `json:"cache_write_tokens"`
			AudioTokens      int `json:"audio_tokens"`
			VideoTokens      int `json:"video_tokens"`
		} `json:"prompt_tokens_details"`
		CostDetails struct {
			UpstreamInferenceCost            float64 `json:"upstream_inference_cost"`
			UpstreamInferencePromptCost      float64 `json:"upstream_inference_prompt_cost"`
			UpstreamInferenceCompletionsCost float64 `json:"upstream_inference_completions_cost"`
		} `json:"cost_details"`
		CompletionTokensDetails struct {
			ReasoningTokens int `json:"reasoning_tokens"`
			ImageTokens     int `json:"image_tokens"`
			AudioTokens     int `json:"audio_tokens"`
		} `json:"completion_tokens_details"`
	} `json:"usage"`
}

type OpenRouter struct {
	apiKey string
}

func (o *OpenRouter) RequestLLM(systemPrompt string, userContexts []Message) (string, error) {
	payload := ChatRequest{
		Model: "deepseek/deepseek-v4-flash",
		Messages: []Message{
			{Role: "system", Content: systemPrompt},
		},
	}

	for _, userContext := range userContexts {
		payload.Messages = append(payload.Messages, userContext)
	}

	// 1. Маршалим в JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	// 2. Создаём io.Reader
	bodyReader := bytes.NewReader(jsonData)

	// 3. Используем в запросе (например, с хелпером DoRequest)
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + o.apiKey,
	}

	res, err := helpers.DoRequest("POST", "https://openrouter.ai/api/v1/chat/completions", bodyReader, headers)

	if err != nil {
		return "", err
	}

	var parsed OpenRouterApiResponse

	err = json.Unmarshal([]byte(res), &parsed)

	if err != nil {
		return "", err
	}

	return parsed.Choices[0].Message.Content, nil
}

func NewOpenRouter(apiKey string) *OpenRouter {
	return &OpenRouter{
		apiKey: apiKey,
	}
}
