package app

import (
	"log"

	"github.com/eduwr/go-rinha-de-backend/routes"
	"github.com/gofiber/fiber/v2"
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

func (a *App) Setup() {
	routes.RegisterRoutes(a.Client)
}

func (a *App) Serve(p string) {
	log.Fatal(a.Client.Listen(p))
}
