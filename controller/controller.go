package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// Handlers

func renderIndex(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "titulo index",
		"User":  User{"Lucas", "29"}})
}

func renderForm(c *fiber.Ctx) error {
	return c.Render("form", fiber.Map{"user": User{"Lucas", "25"}})
}

func getUsers(c *fiber.Ctx) error {
	var user User
	var users []User
	result, err := DB.Query("SELECT Name, DNI FROM MemberTable")
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

func createUser(c *fiber.Ctx) error {
	user := User{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(user)
}
