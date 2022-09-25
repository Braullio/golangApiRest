package googleService

import (
	"cloud.google.com/go/bigquery"
	"context"
	"google.golang.org/api/option"
	"log"
)

type BigQuery struct {
	ProjectId    string
	GoogleIamKey string
	StringSql    string
}

func BuildBigQuerySql(googleIamKey string, projectId string, stringSql string) BigQuery {
	var bigqueryStruct BigQuery

	bigqueryStruct.GoogleIamKey = googleIamKey
	bigqueryStruct.StringSql = stringSql
	bigqueryStruct.ProjectId = projectId

	return bigqueryStruct
}

func SendToBigQuery(bigqueryStruct BigQuery) *bigquery.RowIterator {
	if len(bigqueryStruct.ProjectId) == 0 || len(bigqueryStruct.GoogleIamKey) == 0 || len(bigqueryStruct.StringSql) == 0 {
		log.Fatalf("google.bigquery: Variables not valid")
	}

	ctx := context.Background()

	client, err := bigquery.NewClient(
		ctx,
		bigqueryStruct.ProjectId,
		option.WithCredentialsFile(bigqueryStruct.GoogleIamKey))

	if err != nil {
		log.Fatalf("google.bigquery.NewClient: %v", err)
	}

	defer func(client *bigquery.Client) {
		err := client.Close()
		if err != nil {
		}
	}(client)

	query := client.Query(bigqueryStruct.StringSql)
	rows, err := query.Read(ctx)

	if err != nil {
		log.Fatalf("google.bigquery: %v", err)
	}

	return rows
}
