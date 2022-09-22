package bigquery

import (
	"cloud.google.com/go/bigquery"
	"context"
	"fmt"
	"golangApiRest/models"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"os"
)

func RunSql(sqlString string) {
	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, os.Getenv("BIGQUERY_PROJECT_ID"), option.WithCredentialsFile(os.Getenv("GOOGLE_CLOUD_KEY")))

	if err != nil {
		log.Fatalf("bigquery.NewClient: %v", err)
	}

	defer func(client *bigquery.Client) {
		err := client.Close()
		if err != nil {
		}
	}(client)

	_, err = query(ctx, client, sqlString)

	if err != nil {
		log.Fatalf("bigquery.NewClient: %v", err)
	}
}

func query(ctx context.Context, client *bigquery.Client, sqlString string) (*bigquery.RowIterator, error) {
	query := client.Query(sqlString)
	return query.Read(ctx)
}

func printValues(rows *bigquery.RowIterator) {
	for {
		var vals models.User

		err := rows.Next(&vals)

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatalf("bigquery.NewClient: %v", err)
		}

		fmt.Println(vals)
	}
}
