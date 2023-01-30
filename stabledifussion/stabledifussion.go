package stabledifussion

import (
	"bytes"
	"context"
	_ "embed"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"text/template"
)

//go:embed stablediffusion.template
var stableDifussionTemplate string

type stableDiffusionBodyContent struct {
	Request string
}

type StableDiffusion struct {
}

func NewStableDiffusion() *StableDiffusion {
	return &StableDiffusion{}
}

func (sd StableDiffusion) Process(ctx context.Context, md string) (string, error) {
	engineId := "stable-diffusion-512-v2-0"
	apiHost := "https://api.stability.ai"
	reqUrl := apiHost + "/v1alpha/generation/" + engineId + "/text-to-image"
	outFile := "./textToImage.json"
	tmpl := template.Must(template.New("stabledifussionTemplate").Parse(stableDifussionTemplate))

	var templateCompleted bytes.Buffer
	stableDiffusionBody := stableDiffusionBodyContent{
		Request: md,
	}

	if tmplErr := tmpl.Execute(&templateCompleted, stableDiffusionBody); tmplErr != nil {
		return "", tmplErr
	}

	var data = templateCompleted.Bytes()

	req, _ := http.NewRequest("POST", reqUrl, bytes.NewBuffer(data))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", os.Getenv("APIKEY_STABLEDIFUSSION"))

	// Execute the request & read all the bytes of the response
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	file, _ := os.Create(outFile)
	defer file.Close()
	_, err := file.Write(body)
	if err != nil {
		panic(err)
	}
	return string(body), nil
}

func (ch StableDiffusion) StableDiffusionProccessHandle(c *gin.Context) {
	var stableDifussionRequest StableDifussionRequest
	if err := c.BindJSON(&stableDifussionRequest); err != nil {
		return
	}
	response, err := ch.Process(c, stableDifussionRequest.Message)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, response)
}
