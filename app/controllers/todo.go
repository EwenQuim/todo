package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EwenQuim/todo-app/app/model"
	"github.com/EwenQuim/todo-app/app/query"
	"github.com/EwenQuim/todo-app/app/validator"
)

func (rs TodoResources) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := query.GetAllTodos(rs.Service)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(todos)
}

func (rs TodoResources) GetTodo(w http.ResponseWriter, r *http.Request) {
	uuid := r.PathValue("uuid")

	if !validator.UUID(uuid) {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	todo, err := query.GetTodo(rs.Service, uuid)
	if err != nil {
		http.Error(w, "Error:"+err.Error(), http.StatusBadRequest)
		return
	}

	todo.Items, err = query.GetItemsForList(rs.Service, uuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	todo.Groups = make([]model.Group, 0, len(todo.Items))
	todo.Groups = append(todo.Groups, model.Group{Name: ""})
	for _, item := range todo.Items {
		group, _ := validator.GetGroupAndContent(item.Content)

		var found bool
		for i, g := range todo.Groups {
			if g.Name == group {
				todo.Groups[i].Items = append(todo.Groups[i].Items, item)
				found = true
				break
			}
		}
		if !found {
			todo.Groups = append(todo.Groups, model.Group{Name: group, Items: []model.Item{item}})
		}
	}

	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rs TodoResources) NewTodo(w http.ResponseWriter, r *http.Request) {
	newTodo := model.Todo{
		Title:  r.URL.Query().Get("title"),
		Public: r.URL.Query().Get("public") == "true",
	}

	newTodo, err := query.NewTodo(rs.Service, newTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newTodo)
}

func (rs TodoResources) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	err := query.DeleteTodo(rs.Service, r.PathValue("uuid"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Todo deleted"))
}
