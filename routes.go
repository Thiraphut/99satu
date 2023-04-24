package main

import (
	"99living-go/api"
	"99living-go/service"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {

	app.Get("/album", api.GetAlbums)
	app.Post("/album", service.CreateAlbums)
}