package routes

import (
	"log"

	"github.com/eduwr/go-rinha-de-backend/pessoas"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return c.SendString("Hello, World!")
	})

	app.Get("/pessoas", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return c.SendString("Not Implemented!")
	})

	app.Post("/pessoas", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		p := new(pessoas.Pessoa)

		if err := c.BodyParser(p); err != nil {
			return err
		}

		if err := p.Validate(); err != nil {
			return err
		}

		log.Println(p.Apelido)
		log.Println(p.Nascimento)
		return c.JSON(p)
	})

	app.Get("/pessoas/:id", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return c.SendString("Not Implemented!")
	})

	app.Get("/contagem-pessoas", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return c.SendString("Not Implemented!")
	})
}
