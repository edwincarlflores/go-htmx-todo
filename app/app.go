package app

import (
	"log"

	"github.com/a-h/templ"
	"github.com/edwincarlflores/go-htmx-todo/components"
	"github.com/edwincarlflores/go-htmx-todo/types"
	"github.com/gofiber/fiber/v2"
)

func App() {
	app := fiber.New()

	app.Static("/assets", "./assets")

	todos := []types.Todo{
		{
			ID:    1,
			Title: "First item",
			Done:  false,
			Body:  "This is an item",
		},
		{
			ID:    2,
			Title: "Second item",
			Done:  false,
			Body:  "This is another item",
		},
	}

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString(("OK"))
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		return HTML(c, components.Page("Test"))
	})

	app.Post("/clicked", func(c *fiber.Ctx) error {
		return HTML(c, components.Clicked())
	})

	app.Get("/todos", func(c *fiber.Ctx) error {
		return HTML(c, components.TodoList(todos))
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return HTML(c, components.TodosPage())
	})

	log.Fatal(app.Listen(":4000"))
}

func HTML(c *fiber.Ctx, comp templ.Component) error {
	c.Set("Content-Type", "text/html")
	return comp.Render(c.Context(), c.Response().BodyWriter())
}
