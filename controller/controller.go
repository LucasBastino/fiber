package controller

import (
	"log"

	"github.com/LucasBastino/fiber/database"
	"github.com/LucasBastino/fiber/models"
	"github.com/gofiber/fiber/v2"
)

// Handlers

func RenderIndex(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "titulo index",
		"User":  models.User{"Lucas", "29"}})
}

func RenderForm(c *fiber.Ctx) error {
	return c.Render("form", fiber.Map{"user": models.User{"Lucas", "25"}})
}

func GetUsers(c *fiber.Ctx) error {
	var user models.User
	var users []models.User
	result, err := database.DB.Query("SELECT Name, DNI FROM MemberTable")
	if err != nil {
		log.Println(err)
	}
	for result.Next() {
		err = result.Scan(&user.Name, &user.DNI)
		if err != nil {
			log.Println(err)
		}
		users = append(users, user)
	}

	return c.Render("index", fiber.Map{"users": users})

}

func CreateUser(c *fiber.Ctx) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(user)
}
