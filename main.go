package main

import (
	"bb8backend/chatgpt"
	"bb8backend/stabledifussion"
	"github.com/gin-gonic/gin"
)

func main() {

	chatGPT := chatgpt.NewChatGPT()
	stableDifussion := stabledifussion.NewStableDiffusion()
	r := gin.Default()

	r.POST("/chatgpt", chatGPT.ChatGPTProccessHandle)
	r.POST("/text2image", stableDifussion.StableDiffusionProccessHandle)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
