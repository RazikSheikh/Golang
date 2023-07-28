package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))
	app.Get("/", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{"Client": c.IPs(), "ClientIP": c.Context().RemoteIP()})
	})

	log.Fatal(app.Listen(":3000"))
}
