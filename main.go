package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("HELLO RINHA DE BACKEND")
	app := fiber.New(fiber.Config{
		AppName: "Go! Rinha de Backend",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/pessoas", func(c *fiber.Ctx) error {
		return c.SendString("Not Implemented!")
	})

	app.Post("/pessoas", func(c *fiber.Ctx) error {
		return c.SendString("Not Implemented!")
	})

	app.Get("/pessoas/:id", func(c *fiber.Ctx) error {
		return c.SendString("Not Implemented!")
	})

	app.Get("/contagem-pessoas", func(c *fiber.Ctx) error {
		return c.SendString("Not Implemented!")
	})

	log.Fatal(app.Listen(":3000"))

}
