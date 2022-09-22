package userController

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golangApiRest/models"
	bigqueryService "golangApiRest/services"
	"log"
	"time"
)

var DatasetId = "golangApiRest"
var TableId = "users"

func Show(c *fiber.Ctx) error {
	var user models.User

	id := c.Params("id")

	bigqueryService.RunSql(buildSelectForBigquery(id))

	return c.Status(fiber.StatusOK).JSON(&user)
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
	queryString := fmt.Sprintf(
		`Select * FROM %s.%s WHERE id = "%s"`,
		DatasetId, TableId, id)

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
