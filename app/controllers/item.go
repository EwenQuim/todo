package controllers

import (
	"fmt"

	"github.com/EwenQuim/todo-app/app/model"
	"github.com/EwenQuim/todo-app/app/query"
	"github.com/EwenQuim/todo-app/app/validator"
	"github.com/EwenQuim/todo-app/database"
	"github.com/gofiber/fiber/v2"
)

func NewItem(c *fiber.Ctx, s database.Service) error {
	newItem := model.Item{
		Content:  validator.CleanItem(c.Query("content"), s),
		TodoUUID: c.Params("uuid"),
	}

	newItem, err := query.NewItem(s, newItem)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(409)
	}

	return c.Status(201).JSON(newItem)
}

func DeleteItem(c *fiber.Ctx, s database.Service) error {

	err := query.DeleteItem(s, c.Params("itemid"))
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(500)
	}

	return c.SendStatus(201)
}

func SwitchItem(c *fiber.Ctx, s database.Service) error {

	err := query.SwitchItem(s, c.Params("itemid"))
	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(err)
	}

	return c.SendStatus(201)
}
