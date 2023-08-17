package routes

import (
	"fmt"

	"github.com/eduwr/go-rinha-de-backend/pessoas"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type CustomContext struct {
	*fiber.Ctx
	DB *sqlx.DB
}

func RegisterRoutes(app *fiber.App, db *sqlx.DB) {
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

		p, err := p.Create(db)

		if err != nil {
			return err
		}

		c.Set("Location", fmt.Sprintf("/pessoas/%s", p.Id))
		return c.Status(201).SendString("Pessoa created successfully")
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
