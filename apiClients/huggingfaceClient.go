package apiClients

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	chatModel "portfolioAPI/chat/models"
)

type HuggingFaceClient struct{}

func NewHuggingFaceClient() *HuggingFaceClient {
	return &HuggingFaceClient{}
}

func (client *HuggingFaceClient) GetVectorByText(text string) (chatModel.FeatureExtractionResponse, error) {
	request, err := createRequest(text)

	if err != nil {
		return chatModel.FeatureExtractionResponse{}, err
	}
	setHeaders(request)

	resp, err := getRequestResponse(request)
	if err != nil {
		return chatModel.FeatureExtractionResponse{}, err
	}

	featureExtraction, err := getFeatureExtraction(resp)
	if err != nil {
		return chatModel.FeatureExtractionResponse{}, err
	}

	return featureExtraction, nil
}

func createRequest(text string) (*http.Request, error) {
	hf_inferenceUrl := "https://router.huggingface.co/hf-inference/models/BAAI/bge-small-en-v1.5/pipeline/feature-extraction"
	payload := []byte(fmt.Sprintf(`{
		"inputs": "%s"
	}`, text))

	request, err := http.NewRequest("POST", hf_inferenceUrl, bytes.NewBuffer(payload))

	if err != nil {
		return nil, err
	}

	return request, nil
}

func setHeaders(request *http.Request) {
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+os.Getenv("HF_TOKEN"))
}

func getRequestResponse(request *http.Request) (*http.Response, error) {
	client := &http.Client{}
	res, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return nil, errors.New(string(body))
	}

	return res, nil
}

func getFeatureExtraction(resp *http.Response) (chatModel.FeatureExtractionResponse, error) {
	featureExtractionResponse := chatModel.FeatureExtractionResponse{}

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(resBody, &featureExtractionResponse); err != nil {
		return nil, err
	}

	return featureExtractionResponse, nil
}
