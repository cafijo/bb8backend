package chatgpt

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	gogpt "github.com/sashabaranov/go-gpt3"
	"net/http"
	"os"
)

type ChatGPT struct {
	Client           *gogpt.Client
	MaxToken         int
	Temperature      float32
	TopP             float32
	BestOf           int
	FrequencyPenalty float32
	PresencePenalty  float32
}

func NewChatGPT() *ChatGPT {
	c := gogpt.NewClient(os.Getenv("APIKEY_CHATGPT"))
	chatGPT := ChatGPT{
		Client:           c,
		MaxToken:         500,
		Temperature:      0.7,
		TopP:             1,
		BestOf:           4,
		FrequencyPenalty: 0,
		PresencePenalty:  0.2,
	}
	return &chatGPT
}

func (ch ChatGPT) Process(ctx context.Context, md string) (string, error) {
	req := gogpt.CompletionRequest{
		Model:            gogpt.GPT3TextDavinci003,
		MaxTokens:        ch.MaxToken,
		Temperature:      ch.Temperature,
		TopP:             ch.TopP,
		BestOf:           ch.BestOf,
		FrequencyPenalty: ch.FrequencyPenalty,
		PresencePenalty:  ch.PresencePenalty,
		Prompt:           md,
	}
	resp, err := ch.Client.CreateCompletion(ctx, req)
	if err != nil {
		fmt.Printf("%+v", err)
	}
	return resp.Choices[0].Text, nil
}

func (ch ChatGPT) ChatGPTProccessHandle(c *gin.Context) {
	var chatGPTRequest ChatGPTRequest
	if err := c.BindJSON(&chatGPTRequest); err != nil {
		return
	}
	response, err := ch.Process(c, chatGPTRequest.Message)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": response,
	})
}
