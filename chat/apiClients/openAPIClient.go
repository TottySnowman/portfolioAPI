package apiClients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	chatModel "portfolioAPI/chat/models"
	"strings"
)

type OpenAiClient struct{}

func NewOpenAiClient() *OpenAiClient {
	return &OpenAiClient{}
}

func (apiClient *OpenAiClient) GetSummaryResponse(knowledgeBase []string, prompt string) (string, error) {
	request, err := createOpenAiRequest(knowledgeBase, prompt)
	if err != nil {
		return "", err
	}

	setHeadersOpenAi(request)
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return "", err
	}

  body, err := getOpenAiResponse(response)

	if err != nil {
		return "", err
	}

	return body.Choices[0].Message.Content, nil
}

func createOpenAiRequest(knowledgeBase []string, prompt string) (*http.Request, error) {
	openAiEndpoint := "https://api.openai.com/v1/chat/completions"
	payload := []byte(fmt.Sprintf(`{
    "model": "gpt-4o-mini",
    "store": true,
    "messages": [
        {
            "role": "developer",
            "content": "You are a helpful and friendly assistant to Paul's portfolio. If they ask about you, they probably mean Paul. If you don't know the answer you say: \"Sorry, I am not sure that I can help with that.\" If they ask anything else that doesn't involve Paul or projects created by Paul, you say that you don't know. "
        },
        {
            "role": "user",
    "content": "The question is: %s and the only knowledge you have is this input: %s."
        }
    ]
}`, prompt, strings.ReplaceAll(strings.Join(knowledgeBase, ","), "\"", "\\\"")))

	request, err := http.NewRequest("POST", openAiEndpoint, bytes.NewBuffer(payload))

	if err != nil {
		return nil, err
	}

	return request, nil
}

func setHeadersOpenAi(request *http.Request) {
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+os.Getenv("OPENAI_KEY"))
}

func getOpenAiResponse(response *http.Response) (*chatModel.OpenAIResponse, error) {
	var openAIResponse chatModel.OpenAIResponse

	bodyBytes, err := io.ReadAll(response.Body)

	if err != nil {
    return nil, err
	}

	err = json.Unmarshal(bodyBytes, &openAIResponse)
	if err != nil {
    return nil, err
	}

  return &openAIResponse, nil
}
