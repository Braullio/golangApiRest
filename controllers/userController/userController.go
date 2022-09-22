package userController

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golangApiRest/models"
	bigqueryService "golangApiRest/services"
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

	rows := bigqueryService.RunSql(buildSelectForBigquery(id))

	var users []models.User

	for {
		var user models.User

		err := rows.Next(&user)

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatalf("bigquery.Next: %v", err)
		}

		users = append(users, user)
	}

	return c.Status(fiber.StatusOK).JSON(&users)
}

func Create(c *fiber.Ctx) error {
	var user models.User

	err := json.Unmarshal(c.Body(), &user)

	if err != nil {
		log.Fatalln(err)
	}

	timeNow := time.Now()

	user.Id = uuid.New().String()
	user.Created = timeNow.Format("2006-01-02 15:04:05")
	user.Updated = timeNow.Format("2006-01-02 15:04:05")
	user.CreatedTime = timeNow
	user.UpdatedTime = timeNow

	bigqueryService.RunSql(buildInsertForBigquery(user))

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
	user.Updated = timeNow.Format("2006-01-02 15:04:05")

	bigqueryService.RunSql(buildUpdateForBigquery(id, user))

	return c.SendStatus(fiber.StatusOK)
}

func Delete(c *fiber.Ctx) error {

	id := c.Params("id")

	bigqueryService.RunSql(buildDeleteForBigquery(id))

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

func buildInsertForBigquery(user models.User) string {
	queryString := fmt.Sprintf(
		`INSERT INTO  %s.%s (id, status, name, phone, created_at, updated_at) VALUES ( "%s","%s","%s","%s","%s","%s")`,
		DatasetId, TableId,
		user.Id, user.Status, user.Name, user.Phone, user.Created, user.Updated)

	return queryString
}

func buildUpdateForBigquery(id string, user models.User) string {
	queryString := fmt.Sprintf(
		`UPDATE %s.%s SET status = "%s", updated_at = "%s" WHERE id = "%s"`,
		DatasetId, TableId, user.Status, user.Updated, id)

	return queryString
}

func buildDeleteForBigquery(id string) string {
	queryString := fmt.Sprintf(
		`DELETE FROM %s.%s WHERE id = "%s"`,
		DatasetId, TableId, id)

	return queryString
}
