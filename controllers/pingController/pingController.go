package pingController

import (
	"github.com/gofiber/fiber/v2"
	"golangApiRest/models"
	"golangApiRest/services/googleService"
	"os"
)

func Index(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(&models.Ping{Msg: "pong"})
}

func Teste(c *fiber.Ctx) error {
	chat := os.Getenv("BIGQUERY_PROJECT_ID")
	text := "teste"

	googleService.NotficationInChat(chat, text)

	return c.SendStatus(fiber.StatusOK)
}
