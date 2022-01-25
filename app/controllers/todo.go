package controllers

import (
	"fmt"

	"github.com/EwenQuim/todo-app/app/model"
	"github.com/EwenQuim/todo-app/app/query"
	"github.com/EwenQuim/todo-app/app/validator"
	"github.com/EwenQuim/todo-app/database"
	"github.com/gofiber/fiber/v2"
)

func GetAllTodos(c *fiber.Ctx, s database.Service) error {
	todos, err := query.GetAllTodos(s)
	if err != nil {
		return c.Status(404).SendString(fmt.Sprintf("Error: %s", err.Error()))
	}

	return c.JSON(todos)
}

func GetTodo(c *fiber.Ctx, s database.Service) error {
	uuid := c.Params("uuid")

	if !validator.UUID(uuid) {
		return c.Status(404).SendString(fmt.Sprintf("Error: %s", "Invalid UUID"))
	}

	todo, err := query.GetTodo(s, uuid)
	if err != nil {
		return c.Status(404).SendString(fmt.Sprintf("Error: %s", err.Error()))
	}

	todo.Items, err = query.GetItemsForList(s, uuid)
	if err != nil {
		return c.Status(500).SendString(fmt.Sprintf("Error: %s", err.Error()))
	}

	todo.Groups = make(map[string][]model.Item)
	for _, item := range todo.Items {
		group, _ := validator.GetGroupAndContent(item.Content)
		if len(todo.Groups[group]) == 0 {
			todo.Groups[group] = []model.Item{}
		}
		todo.Groups[group] = append(todo.Groups[group], item)
	}

	return c.JSON(todo)
}

func NewTodo(c *fiber.Ctx, s database.Service) error {
	newTodo := model.Todo{
		Title:  c.Query("title"),
		Public: c.Query("public") == "true",
	}

	newTodo, err := query.NewTodo(s, newTodo)
	if err != nil {
		fmt.Println(err)
		return c.Status(409).JSON(err)
	}

	return c.Status(201).JSON(newTodo)
}

func DeleteTodo(c *fiber.Ctx, s database.Service) error {
	err := query.DeteleTodo(s, c.Params("uuid"))
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(500)
	}

	return c.SendStatus(201)
}
