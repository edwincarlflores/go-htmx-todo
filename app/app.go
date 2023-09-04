package app

import (
	"log"

	"github.com/a-h/templ"
	"github.com/edwincarlflores/go-htmx-todo/components"
	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

func App() {
	app := fiber.New()

	app.Static("/assets", "./assets")

	todos := []Todo{}

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString(("OK"))
	})

	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		return HTML(c, components.Page("Test", "John"))
	})

	app.Get("/tested", func(c *fiber.Ctx) error {
		return HTML(c, components.Page("Tested", "Ed"))
	})

	app.Post("/clicked", func(c *fiber.Ctx) error {
		return HTML(c, components.Clicked())
	})

	log.Fatal(app.Listen(":4000"))
}

func HTML(c *fiber.Ctx, comp templ.Component) error {
	c.Set("Content-Type", "text/html")
	return comp.Render(c.Context(), c.Response().BodyWriter())
}
