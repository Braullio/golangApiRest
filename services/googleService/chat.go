package googleService

import (
	"fmt"
	"golangApiRest/services/enumChat"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type Chat struct {
	Message string
	Webhook string
}

const (
	maxLength = 3980
	codeBlock = "```"
)

func BuildChatAlertMessage(tag enumChat.ChatType, webhook string, functionExecution string, message string, observation string, attempt int) Chat {
	var chatStruct Chat

	chatStruct.Webhook = webhook

	if len(functionExecution) != 0 && len(message) != 0 {
		timeNow := time.Now()

		text := fmt.Sprintf(
			`{"text": "%s\n\n*DateTime:* _%s_\n*Attempts:* _%d_\n*Function:* _%s_\n*Message:*\n%s%s`,
			tag, timeNow.Format("2006-01-02 15:04:05"), attempt, functionExecution, codeBlock, message,
		)

		if len(observation) > 0 {
			text += fmt.Sprintf("%s*Observation:*\n%s%s", codeBlock, codeBlock, observation)
		}

		if len(text) > maxLength {
			rs := []rune(text)
			text = string(rs[:maxLength])
		}

		chatStruct.Message = fmt.Sprintf(`%s ...%s"}`, text, codeBlock)
	}

	return chatStruct
}

func SendToChat(chatStruct Chat) {
	if len(chatStruct.Webhook) == 0 {
		log.Printf("[SendToChat] google.chat: Variables not valid")

		return
	}

	payload := strings.NewReader(chatStruct.Message)

	client := &http.Client{}
	res, err := client.Post(chatStruct.Webhook, "application/json", payload)

	if err != nil {
		log.Printf("[SendToChat] google.chat: %v", err)

		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(res.Body)
}
