package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID    int    `json:"id"`
	TITLE string `json:"title"`

	Done bool `json:"done"`

	Body string `json:"body"`
}

func main() {
	fmt.Println("hello")
	todos := []Todo{}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}
		if err := c.BodyParser(todo); err != nil {
			return err
		}
		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.JSON(todos)

	})

	app.Patch("api/todos/:id/done", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt(("id"))
		if err != nil {
			return c.Status(401).SendString("Invalid id")
		}
		for i, todo := range todos {
			if todo.ID == id {
				todos[i].Done = true
			}
			break
		}
		return c.Status(201).JSON(todos)

	})
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	log.Fatal(app.Listen(":5000"))
}
