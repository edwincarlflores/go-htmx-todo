package main

import (
	"log"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

func main() {
	component := hello("John")
	app := fiber.New()

	todos := []Todo{}

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString(("OK"))
	})

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	app.Get("/api/test", func(c *fiber.Ctx) error {
		return HTML(c, component)
	})

	app.Get("/api/tested", func(c *fiber.Ctx) error {
		return HTML(c, hello("Ed"))
	})

	log.Fatal(app.Listen(":4000"))
}

func HTML(c *fiber.Ctx, comp templ.Component) error {
	c.Set("Content-Type", "text/html")
	return comp.Render(c.Context(), c.Response().BodyWriter())
}
