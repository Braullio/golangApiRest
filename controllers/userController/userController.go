package userController

import (
	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/civil"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golangApiRest/models"
	"golangApiRest/services/enumChat"
	"golangApiRest/services/googleService"
	"google.golang.org/api/iterator"
	"log"
	"os"
	"time"
)

var DatasetId = "golangApiRest"
var TableId = "users"
var attempt = 0

func getEnv() (string, string, string) {
	ProjectId := os.Getenv("BIGQUERY_PROJECT_ID")
	GoogleIamKey := os.Getenv("GOOGLE_IAM_KEY")
	GoogleChatWebhook := os.Getenv("GOOGLE_CHAT_WEBHOOK")

	return ProjectId, GoogleIamKey, GoogleChatWebhook
}

func Show(c *fiber.Ctx) error {
	ProjectId, GoogleIamKey, _ := getEnv()

	var id string

	if len(c.Params("id")) > 0 {
		id = c.Params("id")
	} else {
		id = ""
	}

	usersStruct := buildUsersToResponse(
		googleService.SendToBigQuery(
			GoogleIamKey,
			ProjectId,
			buildSelectForBigquery(id),
			attempt,
		),
	)

	return c.Status(fiber.StatusOK).JSON(&usersStruct)
}

func Create(c *fiber.Ctx) error {
	ProjectId, GoogleIamKey, GoogleChatWebhook := getEnv()

	var user models.User

	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		go googleService.SendToChat(
			googleService.BuildChatAlertMessage(
				enumChat.Warning,
				GoogleChatWebhook,
				"userController.Create[ json.Unmarshal ]",
				err.Error(),
				"",
				attempt,
			),
		)

		return c.SendStatus(fiber.StatusBadRequest)
	}

	timeNow := time.Now()

	user.Id = uuid.New().String()
	user.Created = civil.DateTimeOf(timeNow)
	user.Updated = civil.DateTimeOf(timeNow)

	go googleService.SendToBigQuery(
		GoogleIamKey,
		ProjectId,
		buildInsertForBigquery(user, timeNow),
		attempt,
	)

	return c.Status(fiber.StatusCreated).JSON(&user)
}

func Update(c *fiber.Ctx) error {
	ProjectId, GoogleIamKey, GoogleChatWebhook := getEnv()

	var user models.User

	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		go googleService.SendToChat(
			googleService.BuildChatAlertMessage(
				enumChat.Warning,
				GoogleChatWebhook,
				"userController.Update[ json.Unmarshal ]",
				err.Error(),
				"",
				attempt,
			),
		)

		return c.SendStatus(fiber.StatusBadRequest)
	}

	timeNow := time.Now()
	user.Id = c.Params("id")
	if len(user.Id) == 0 {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	go googleService.SendToBigQuery(
		GoogleIamKey,
		ProjectId,
		buildUpdateForBigquery(user, timeNow),
		attempt,
	)

	return c.SendStatus(fiber.StatusOK)
}

func Delete(c *fiber.Ctx) error {
	ProjectId, GoogleIamKey, GoogleChatWebhook := getEnv()

	id := c.Params("id")
	if len(id) == 0 {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	go googleService.SendToBigQuery(
		GoogleIamKey,
		ProjectId,
		buildDeleteForBigquery(id),
		attempt,
	)

	go googleService.SendToChat(
		googleService.BuildChatAlertMessage(
			enumChat.Warning,
			GoogleChatWebhook,
			"userController.Update[ json.Unmarshal ]",
			fmt.Sprintf("[GoLangApiRest] ALERT\nEfetuado a deleção do id: %s", id),
			"",
			attempt,
		),
	)

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
