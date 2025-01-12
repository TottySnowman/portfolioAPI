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

type HuggingFaceClient struct{

}

func NewHuggingFaceClient() *HuggingFaceClient{
  return &HuggingFaceClient{}
}

func (client *HuggingFaceClient) GetVectorByText(text string)(chatModel.FeatureExtractionResponse, error){
  request, err := createRequest(text)

  if err != nil{
    return chatModel.FeatureExtractionResponse{}, err
  }
  setHeaders(request)

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with status: %d\n", res.StatusCode)
		body, _ := io.ReadAll(res.Body)
		fmt.Printf("Error response: %s\n", body)
		return nil, errors.New("Failed to get the thing")
	}

	featureExtractionResponse := chatModel.FeatureExtractionResponse{}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(resBody, &featureExtractionResponse); err != nil {
		return nil, err
	}

	return featureExtractionResponse, nil
}

func createRequest(text string)(*http.Request, error){
	hf_inferenceUrl := "https://api-inference.huggingface.co/models/BAAI/bge-small-en-v1.5"
	payload := []byte(fmt.Sprintf(`{
		"inputs": "%s"
	}`, text))

	request, err := http.NewRequest("POST", hf_inferenceUrl, bytes.NewBuffer(payload))

	if err != nil {
    return nil, err
	}

  return request, nil
}

func setHeaders(request *http.Request){
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+os.Getenv("HF_TOKEN"))
}

func doRequest(request *http.Request){

}
