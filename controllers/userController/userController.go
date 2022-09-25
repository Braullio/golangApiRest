package userController

import (
	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/civil"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golangApiRest/models"
	"golangApiRest/services/googleService"
	"google.golang.org/api/iterator"
	"log"
	"os"
	"time"
)

var GoogleChatWebhook = os.Getenv("GOOGLE_CHAT_WEBHOOK")
var GoogleIamKey = os.Getenv("GOOGLE_IAM_KEY")
var ProjectId = os.Getenv("BIGQUERY_PROJECT_ID")
var DatasetId = "golangApiRest"
var TableId = "users"

func Show(c *fiber.Ctx) error {
	var id string

	if len(c.Params("id")) > 0 {
		id = c.Params("id")
	} else {
		id = ""
	}

	bigqueryStruct := googleService.BuildBigQuerySql(GoogleIamKey, ProjectId, buildSelectForBigquery(id))
	usersStruct := buildUsersToResponse(googleService.SendToBigQuery(bigqueryStruct))

	return c.Status(fiber.StatusOK).JSON(&usersStruct)
}

func Create(c *fiber.Ctx) error {
	var user models.User

	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		log.Fatalln(err)
	}

	timeNow := time.Now()

	user.Id = uuid.New().String()
	user.Created = civil.DateTimeOf(timeNow)
	user.Updated = civil.DateTimeOf(timeNow)

	bigqueryStruct := googleService.BuildBigQuerySql(GoogleIamKey, ProjectId, buildInsertForBigquery(user, timeNow))
	googleService.SendToBigQuery(bigqueryStruct)

	return c.Status(fiber.StatusCreated).JSON(&user)
}

func Update(c *fiber.Ctx) error {
	var user models.User

	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		log.Fatalln(err)
	}

	timeNow := time.Now()

	user.Id = c.Params("id")
	user.Updated = civil.DateTimeOf(timeNow)

	bigqueryStruct := googleService.BuildBigQuerySql(GoogleIamKey, ProjectId, buildUpdateForBigquery(user, timeNow))
	googleService.SendToBigQuery(bigqueryStruct)

	return c.SendStatus(fiber.StatusOK)
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	bigqueryStruct := googleService.BuildBigQuerySql(GoogleIamKey, ProjectId, buildDeleteForBigquery(id))
	googleService.SendToBigQuery(bigqueryStruct)

	chatStruct := googleService.BuildChatSimpleMessage(
		GoogleChatWebhook,
		fmt.Sprintf("[GoLangApiRest] Efetuado a deleção do id: %s", id),
	)

	googleService.SendToChat(chatStruct)

	return c.SendStatus(fiber.StatusOK)
}

func buildUsersToResponse(response *bigquery.RowIterator) []models.User {
	var users []models.User
	row := make(map[string]bigquery.Value)

	for {
		err := response.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("bigquery.Next: %v", err)
		}

		var user models.User
		user.Id = row["id"].(string)
		user.Name = row["name"].(string)
		user.Phone = row["phone"].(string)
		user.Status = row["status"].(string)
		user.Created = row["updated_at"].(civil.DateTime)
		user.Updated = row["created_at"].(civil.DateTime)

		users = append(users, user)
	}

	return users
}

func buildSelectForBigquery(id string) string {
	var queryString string

	if len(id) > 0 {
		queryString = fmt.Sprintf(
			`Select * FROM %s.%s WHERE id = "%s"`,
			DatasetId, TableId, id)
	} else {
		queryString = fmt.Sprintf(
			`Select * FROM %s.%s WHERE id is not null`,
			DatasetId, TableId)
	}

	return queryString
}

func buildInsertForBigquery(user models.User, timeNow time.Time) string {
	queryString := fmt.Sprintf(
		`INSERT INTO  %s.%s (id, status, name, phone, created_at, updated_at) VALUES ( "%s","%s","%s","%s","%s","%s")`,
		DatasetId, TableId,
		user.Id, user.Status, user.Name, user.Phone, timeNow.Format("2006-01-02 15:04:05"), timeNow.Format("2006-01-02 15:04:05"))

	return queryString
}

func buildUpdateForBigquery(user models.User, timeNow time.Time) string {
	queryString := fmt.Sprintf(
		`UPDATE %s.%s SET status = "%s", updated_at = "%s" WHERE id = "%s"`,
		DatasetId, TableId, user.Status, timeNow.Format("2006-01-02 15:04:05"), user.Id)

	return queryString
}

func buildDeleteForBigquery(id string) string {
	queryString := fmt.Sprintf(
		`DELETE FROM %s.%s WHERE id = "%s"`,
		DatasetId, TableId, id)

	return queryString
}
