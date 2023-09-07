package app

import (
	"log"
	"slices"

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
		return HTML(c, components.TestPage("Prime"))
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

	app.Post("/todos/toggle/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		for idx, todo := range todos {
			if todo.ID == id {
				todos[idx].Done = !todos[idx].Done
				return HTML(c, components.TodoItem(todos[idx]))
			}
		}

		return nil
	})

	app.Delete("/todos/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		todos = slices.DeleteFunc(todos, func(todo types.Todo) bool {
			return id == todo.ID
		})

		return nil

	})

	log.Fatal(app.Listen(":8080"))
}

func HTML(c *fiber.Ctx, comp templ.Component) error {
	c.Set("Content-Type", "text/html")
	return comp.Render(c.Context(), c.Response().BodyWriter())
}
