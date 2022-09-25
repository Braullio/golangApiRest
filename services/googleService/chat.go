package googleService

import (
	"fmt"
	"golangApiRest/models"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func NotficationInChat(chat models.Chat) int {
	if len(chat.Webhook) == 0 {
		chat.Webhook = os.Getenv("GOOGLE_CHAT_WEBHOOK")
	}

	if len(chat.Webhook) == 0 {
		return 500
	}

	payload := strings.NewReader(fmt.Sprintf(`{"text": "%s"}`, chat.Text))

	client := &http.Client{}
	res, err := client.Post(chat.Webhook, "application/json", payload)

	if err != nil {
		log.Fatalf("google.chat.NewClient: %v", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(res.Body)

	return res.StatusCode
}
