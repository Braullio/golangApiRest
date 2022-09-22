package main

import (
	"github.com/gofiber/fiber/v2"
	"golangApiRest/controllers/pingController"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/ping", pingController.Index)

	err := app.Listen(":8001")

	if err != nil {
		log.Fatal(err)
	}
}
