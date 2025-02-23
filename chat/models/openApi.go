package chatModel
type OpenAIResponse struct {
    ID      string `json:"id"`
    Choices []struct {
        Message struct {
            Content string `json:"content"`
        } `json:"message"`
    } `json:"choices"`
}
