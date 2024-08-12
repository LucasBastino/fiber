package main

import (
	"log"

	"github.com/LucasBastino/fiber/database"
	"github.com/LucasBastino/fiber/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func main() {

	// app Config
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	// Database connection
	database.CreateConnection()

	// Middlewares
	app.Use(logger.New())

	// Serve static files
	app.Static("/static", "./static")

	// Routing
	routes.SetupRoutes(app)

	// Start
	log.Fatal(app.Listen(":8085"))

}
