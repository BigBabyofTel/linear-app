package api

import (
	"context"
	"net/http"

	openai "github.com/sashabaranov/go-openai"
)

func (a *app) autocompleteHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Text string `json:"text"`
	}

	if err := a.readJSON(w, r, &input); err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	resp, err := a.AI.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       openai.GPT3Dot5Turbo,
			MaxTokens:   40,
			Temperature: 0.5,
			TopP:        1.0,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Finish the user's sentence:",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: input.Text,
				},
			},
		},
	)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	if err := a.writeJSON(w, http.StatusOK, envelope{"data": resp}, nil); err != nil {
		a.serverErrorResponse(w, r, err)
	}
}
