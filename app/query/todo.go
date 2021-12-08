package query

import (
	"github.com/EwenQuim/todo-app/app/model"
	"github.com/EwenQuim/todo-app/database"
	"github.com/google/uuid"
)

func NewTodo(s database.Service, todo model.Todo) (model.Todo, error) {
	uuid := uuid.NewString()
	todo.UUID = uuid

	err := s.DB.Create(&todo).Error
	return todo, err
}

func GetTodo(s database.Service, uuid string) (model.Todo, error) {
	var todo model.Todo
	err := s.DB.First(&todo, "uuid = ?", uuid).Error
	return todo, err
}

func GetAllTodos(s database.Service) ([]model.Todo, error) {
	var todos []model.Todo
	err := s.DB.Where("public = ?", true).Find(&todos).Error
	return todos, err
}

func DeteleTodo(s database.Service, uuid string) error {
	return s.DB.Delete(&model.Todo{}, "uuid = ?", uuid).Error
}
