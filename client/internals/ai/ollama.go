package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type OllamaCLient struct {
	BaseURL string
	Model string
}

type Message struct {
	Role string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model string `json:"model"`
	Prompt string `json:"promt"`
	Stream bool `json:"stream"`
}

type ChatResponse struct {
	Response string `json:"response"`
}

func NewOllama(baseURL, model string) *OllamaCLient {
	return &OllamaCLient{
		BaseURL: baseURL,
		Model: model,
	}
}

func (o *OllamaCLient) Chat(promt string) (string, error) {
	reqBody := ChatRequest{
		Model: o.Model,
		Prompt: promt,
		Stream: false,
	}

	data, _ := json.Marshal(reqBody)

	resp, err := http.Post(
		fmt.Sprintf("%s/api/generate", o.BaseURL),
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Response, nil

}


