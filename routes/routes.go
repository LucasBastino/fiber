package routes

import (
	c "github.com/LucasBastino/fiber/controller"
	"github.com/gofiber/fiber/v2"
)

// Routes
func SetupRoutes(app *fiber.App) {
	app.Get("/", c.RenderIndex)
	app.Get("/form", c.RenderForm)
	app.Get("/users", c.GetUsers)
	app.Post("/createUser", c.CreateUser)
}
