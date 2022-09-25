package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golangApiRest/controllers/pingController"
	"golangApiRest/controllers/productController"
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

	app.Get("/user", userController.Show)
	app.Get("/user/:id", userController.Show)
	app.Post("/user", userController.Create)
	app.Put("/user/:id", userController.Update)
	app.Delete("/user/:id", userController.Delete)

	app.Get("/product", productController.Show)
	app.Get("/product/:id", productController.Show)
	app.Post("/product", productController.Create)
	app.Put("/product/:id", productController.Update)
	app.Delete("/product/:id", productController.Delete)

	err := app.Listen(":8001")

	if err != nil {
		log.Fatal(err)
	}
}
