package routes

import (
	"fmt"
	"strconv"

	"github.com/eduwr/go-rinha-de-backend/pessoas"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type CustomContext struct {
	*fiber.Ctx
	DB *sqlx.DB
}

func RegisterRoutes(app *fiber.App, db *sqlx.DB) {
	app.Get("/pessoas", func(c *fiber.Ctx) error {
		c.Accepts("application/json")

		t := c.Query("t")
		p, err := pessoas.Index(t, db)
		if err != nil {
			return c.Status(422).SendString(err.Error())
		}

		return c.Status(200).JSON(p)
	})

	app.Get("/pessoas/:id", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		id := c.Params("id")

		p, err := pessoas.Show(id, db)

		if err != nil {
			return c.Status(404).SendString("Not Found")
		}

		return c.Status(200).JSON(p)
	})

	app.Post("/pessoas", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		p := pessoas.PessoaWithStack{}

		if err := c.BodyParser(&p); err != nil {
			return c.Status(422).SendString("Unprocessable Entity")
		}

		createdP, err := pessoas.Create(p, db)

		if err != nil {
			return c.Status(422).SendString(err.Error())
		}

		c.Set("Location", fmt.Sprintf("/pessoas/%s", createdP.Id))
		return c.Status(201).SendString("pessoa created successfully")
	})

	app.Get("/contagem-pessoas", func(c *fiber.Ctx) error {
		c.Accepts("application/json")

		count := pessoas.Count(db)
		return c.Status(200).SendString(strconv.Itoa(count))
	})
}
