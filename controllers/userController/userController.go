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
	"time"
)

var DatasetId = "golangApiRest"
var TableId = "users"

func Show(c *fiber.Ctx) error {
	var id string

	if len(c.Params("id")) > 0 {
		id = c.Params("id")
	} else {
		id = ""
	}

	rows := googleService.RunSqlInBigQuery(buildSelectForBigquery(id))

	var users []models.User

	row := make(map[string]bigquery.Value)

	for {
		err := rows.Next(&row)

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatalf("bigquery.Next: %v", err)
		}

		setAtributesInUser(row, &users)
	}

	return c.Status(fiber.StatusOK).JSON(&users)
}

func setAtributesInUser(values map[string]bigquery.Value, users *[]models.User) {
	var user models.User
	user.Id = values["id"].(string)
	user.Name = values["name"].(string)
	user.Phone = values["phone"].(string)
	user.Status = values["status"].(string)
	user.Created = values["updated_at"].(civil.DateTime)
	user.Updated = values["created_at"].(civil.DateTime)

	*users = append(*users, user)
}

func Create(c *fiber.Ctx) error {
	var user models.User

	err := json.Unmarshal(c.Body(), &user)

	if err != nil {
		log.Fatalln(err)
	}

	user.Id = uuid.New().String()

	timeNow := time.Now()
	user.Created = civil.DateTimeOf(timeNow)
	user.Updated = civil.DateTimeOf(timeNow)

	googleService.RunSqlInBigQuery(buildInsertForBigquery(user, timeNow))

	return c.Status(fiber.StatusCreated).JSON(&user)
}

func Update(c *fiber.Ctx) error {
	var user models.User

	err := json.Unmarshal(c.Body(), &user)

	if err != nil {
		log.Fatalln(err)
	}

	id := c.Params("id")

	timeNow := time.Now()
	user.Updated = civil.DateTimeOf(timeNow)

	googleService.RunSqlInBigQuery(buildUpdateForBigquery(id, user, timeNow))

	return c.SendStatus(fiber.StatusOK)
}

func Delete(c *fiber.Ctx) error {

	id := c.Params("id")

	googleService.RunSqlInBigQuery(buildDeleteForBigquery(id))

	return c.SendStatus(fiber.StatusOK)
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

func buildUpdateForBigquery(id string, user models.User, timeNow time.Time) string {
	queryString := fmt.Sprintf(
		`UPDATE %s.%s SET status = "%s", updated_at = "%s" WHERE id = "%s"`,
		DatasetId, TableId, user.Status, timeNow.Format("2006-01-02 15:04:05"), id)

	return queryString
}

func buildDeleteForBigquery(id string) string {
	queryString := fmt.Sprintf(
		`DELETE FROM %s.%s WHERE id = "%s"`,
		DatasetId, TableId, id)

	return queryString
}
