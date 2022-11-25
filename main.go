package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
)

type User struct {
	Id string
	FirstName string
	LastName  string
}

func handleUser (c *fiber.Ctx) error {
	user := User{
		FirstName: "John",
		LastName: "Doe",
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func handleCreateUser (c *fiber.Ctx) error {
	user := User{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.Id = uuid.New().String()

	return c.Status(fiber.StatusCreated).JSON(user)
}

func main() {
	app := fiber.New()

	//Middlewares

	//Ver peticiones en consola
	app.Use(logger.New())

	app.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	userGroup := app.Group("/users")
	userGroup.Get("", handleUser)
	userGroup.Post("", handleCreateUser)

	app.Listen(":3000")
}