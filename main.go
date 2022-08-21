package main

import (
	"fiber-gorm-tutorial/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome API!")
}

func main() {
	database.ConnectDb()

	app := fiber.New()

	app.Get("/", welcome)

	log.Fatal(app.Listen(":3000"))
}
