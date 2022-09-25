package googleService

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Chat struct {
	Message string
	Webhook string
}

func BuildChatSimpleMessage(webhook string, message string) Chat {
	var chatStruct Chat

	chatStruct.Webhook = webhook
	chatStruct.Message = fmt.Sprintf(`{"text": "%s"}`, message)

	return chatStruct
}

func SendToChat(chatStruct Chat) int {
	if len(chatStruct.Webhook) == 0 || len(chatStruct.Message) == 0 {
		log.Fatalf("google.chat: Variables not valid")
	}

	payload := strings.NewReader(chatStruct.Message)

	client := &http.Client{}
	res, err := client.Post(chatStruct.Webhook, "application/json", payload)

	if err != nil {
		log.Fatalf("google.chat: %v", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(res.Body)

	return res.StatusCode
}
