package main

import (
	"log"
	"net"

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
		clientIp := c.IP()

		names, err := net.LookupIP(clientIp)

		if err != nil {
			return c.JSON(fiber.Map{"error": err})
		}

		return c.JSON(fiber.Map{"Client": names})
	})

	log.Fatal(app.Listen(":3000"))
}
