package main

import "github.com/gofiber/fiber/v2"

// Routes
func SetupRoutes(app *fiber.App) {
	app.Get("/", renderIndex)
	app.Get("/form", renderForm)
	app.Get("/users", getUsers)
	app.Post("/createUser", createUser)
}
