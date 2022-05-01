package main

import (
	"github.com/gofiber/fiber"
	"my-rest-api/handlers"
)

const port = 8000

func main() {
	app := fiber.New()
	app.Get("/person/find", handlers.FindPerson)
	app.Get("/person/:id?", handlers.GetPerson)
	app.Post("/person", handlers.CreatePerson)
	app.Put("/person/:id", handlers.UpdatePerson)
	app.Delete("/person/:id", handlers.DeletePerson)

	err := app.Listen(port)
	if err != nil {
		return
	}
}
