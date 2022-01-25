package controllers

import (
	"github.com/EwenQuim/todo-app/database"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, s database.Service) {
	api := app.Group("/api")

	// Todo
	api.Get("/todo/new", wrapService(s, NewTodo))
	api.Get("/todo", wrapService(s, GetAllTodos))
	api.Get("/todo/:uuid", wrapService(s, GetTodo))
	api.Get("/todo/:uuid/delete", wrapService(s, DeleteTodo))

	// Items
	api.Get("/todo/:uuid/new", wrapService(s, NewItem))
	api.Get("/todo/:uuid/:itemid/delete", wrapService(s, DeleteItem))
	api.Get("/todo/:uuid/:itemid/switch", wrapService(s, SwitchItem))

	// Other
	api.Get("/ping", wrapService(s, ping))
}

func wrapService(s database.Service, f func(c *fiber.Ctx, s database.Service) error) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return f(c, s)
	}
}

func ping(c *fiber.Ctx, s database.Service) error {
	return c.SendString("pong")
}
