package googleService

import (
	"cloud.google.com/go/bigquery"
	"context"
	"google.golang.org/api/option"
	"log"
)

func SendToBigQuery(googleIamKey string, projectId string, stringSql string) *bigquery.RowIterator {
	if len(projectId) == 0 || len(googleIamKey) == 0 || len(stringSql) == 0 {
		log.Fatalf("google.bigquery: Variables not valid")
	}

	ctx := context.Background()

	client, err := bigquery.NewClient(
		ctx,
		projectId,
		option.WithCredentialsFile(googleIamKey))

	if err != nil {
		log.Fatalf("google.bigquery.NewClient: %v", err)
	}

	defer func(client *bigquery.Client) {
		err := client.Close()
		if err != nil {
		}
	}(client)

	query := client.Query(stringSql)
	rows, err := query.Read(ctx)

	if err != nil {
		log.Fatalf("google.bigquery: %v", err)
	}

	return rows
}
