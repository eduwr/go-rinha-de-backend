package app

import (
	"log"

	"github.com/eduwr/go-rinha-de-backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type App struct {
	AppName string
	Client  *fiber.App
}

func NewApp(n string) App {
	app := fiber.New(fiber.Config{
		AppName: n,
	})

	return App{
		AppName: n,
		Client:  app,
	}
}

func (a *App) Setup(db *sqlx.DB) {
	routes.RegisterRoutes(a.Client, db)
}

func (a *App) Serve(p string) {
	log.Fatal(a.Client.Listen(p))
}
