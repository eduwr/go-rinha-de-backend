package main

import (
	"fmt"
	"log"

	"github.com/eduwr/go-rinha-de-backend/pessoas"
	"github.com/gofiber/fiber/v2"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("HELLO RINHA DE BACKEND")

	db, err := sqlx.Connect("postgres", "host=localhost port=5432 dbname=rinha user=user password=pass sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(pessoas.PessoaSchema)

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

	log.Fatal(app.Listen(":3333"))

}
