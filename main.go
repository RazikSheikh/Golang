package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{"Client": c.Context().RemoteAddr()})
	})

	log.Fatal(app.Listen(":3000"))
}
