package googleService

import (
	"cloud.google.com/go/bigquery"
	"context"
	"google.golang.org/api/option"
	"log"
	"os"
)

func RunSqlInBigQuery(sqlString string) *bigquery.RowIterator {
	ctx := context.Background()

	client, err := bigquery.NewClient(
		ctx,
		os.Getenv("BIGQUERY_PROJECT_ID"),
		option.WithCredentialsFile(os.Getenv("GOOGLE_CLOUD_KEY")))

	if err != nil {
		log.Fatalf("bigquery.NewClient: %v", err)
	}

	defer func(client *bigquery.Client) {
		err := client.Close()
		if err != nil {
		}
	}(client)

	query := client.Query(sqlString)
	rows, err := query.Read(ctx)

	if err != nil {
		log.Fatalf("bigquery: %v", err)
	}

	return rows
}
