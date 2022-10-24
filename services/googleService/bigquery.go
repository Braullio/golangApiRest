package googleService

import (
	"cloud.google.com/go/bigquery"
	"context"
	"golangApiRest/services/enumChat"
	"google.golang.org/api/option"
	"os"
)

func SendToBigQuery(googleIamKey string, projectId string, stringSql string, attempt int) *bigquery.RowIterator {
	var googleChatWebhook = os.Getenv("GOOGLE_CHAT_WEBHOOK")

	if len(projectId) == 0 || len(googleIamKey) == 0 || len(stringSql) == 0 {
		go SendToChat(
			BuildChatAlertMessage(
				enumChat.Danger,
				googleChatWebhook,
				"SendToBigQuery[ initial validation ]",
				"Variables not valid ",
				stringSql,
				attempt,
			),
		)
	}

	ctx := context.Background()

	client, err := bigquery.NewClient(
		ctx,
		projectId,
		option.WithCredentialsFile(googleIamKey),
	)

	if err != nil {
		go SendToChat(
			BuildChatAlertMessage(
				enumChat.Warning,
				googleChatWebhook,
				"SendToBigQuery[ bigquery.NewClient ]",
				err.Error(),
				stringSql,
				attempt,
			),
		)
	}

	defer func(client *bigquery.Client) {
		err := client.Close()
		if err != nil {
		}
	}(client)

	query := client.Query(stringSql)
	rows, err := query.Read(ctx)

	if err != nil {
		go SendToChat(
			BuildChatAlertMessage(
				enumChat.Warning,
				googleChatWebhook,
				"SendToBigQuery[ query.Read ]",
				err.Error(),
				stringSql,
				attempt,
			),
		)
	}

	return rows
}
