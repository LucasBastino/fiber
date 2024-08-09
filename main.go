package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func main() {

	// app Config
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	db = CreateConnection()
	defer db.Close()
	// Middlewares
	app.Use(logger.New())
	SetupRoutes(app)

	// Start
	log.Fatal(app.Listen(":8085"))

}
