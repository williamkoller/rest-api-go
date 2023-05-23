package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamkoller/rest-api-go/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/facts", handlers.FindFact)

	app.Post("/facts", handlers.CreateFact)
}
