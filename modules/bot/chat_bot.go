package bot

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gtkit/go-openai"
	"github.com/labstack/echo/v4"
)

type Input struct {
	Question string `json:"question"`
}

type Response struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

func ClassifyEnvironmentalIssue(c echo.Context) error {
	// Parse request body
	var input Input
	if err := c.Bind(&input); err != nil {
		return err
	}

	// Initialize OpenAI client
	//apikey := os.Getenv("OPENAI_API_KEY")
	apiKey := "ak-kqwdnpqwdkjmpo"
	if apiKey == "" {
		return c.JSON(http.StatusInternalServerError, Response{Status: "error", Data: "Please set your OpenAI API key."})
	}
	client := openai.NewClient(apiKey)

	// Generate response using OpenAI
	//prompt := fmt.Sprintf("Sebagai pakar lingkungan, saya akan membantu mengklasifikasikan masalah pencemaran lingkungan yang Anda laporkan. Tolong ceritakan masalah atau kejadian yang Anda alami terkait lingkungan hidup. Masalah atau kejadian: %s", input.Question)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{Role: "system", Content: fmt.Sprintf("Sebagai pakar lingkungan, saya akan membantu mengklasifikasikan masalah pencemaran lingkungan yang Anda laporkan. Tolong ceritakan masalah atau kejadian yang Anda alami terkait lingkungan hidup. Masalah atau kejadian: %s", input.Question)},
			},
		},
	)
	if err != nil {
		errResponse := Response{
			Status: "error",
			Data:   err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, errResponse)
	}

	// Extract response from completion
	response := resp.Choices[0].Message.Content

	// Create response
	responseObj := Response{
		Status: "success",
		Data:   response,
	}

	return c.JSON(http.StatusOK, responseObj)
}
