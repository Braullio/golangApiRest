package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golangApiRest/controllers/pingController"
	"golangApiRest/controllers/userController"
	"log"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	app := fiber.New()

	app.Get("/ping", pingController.Index)
	app.Get("/teste", pingController.Teste)

	app.Get("/user", userController.Show)
	app.Get("/user/:id", userController.Show)

	app.Post("/user", userController.Create)
	app.Put("/user/:id", userController.Update)
	app.Delete("/user/:id", userController.Delete)

	err := app.Listen(":8001")

	if err != nil {
		log.Fatal(err)
	}
}
