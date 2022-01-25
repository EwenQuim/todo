package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EwenQuim/todo-app/app/model"
	"github.com/EwenQuim/todo-app/app/query"
	"github.com/EwenQuim/todo-app/app/validator"
	"github.com/go-chi/chi/v5"
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
	uuid := chi.URLParam(r, "uuid")

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

	todo.Groups = make(map[string][]model.Item)
	for _, item := range todo.Items {
		group, _ := validator.GetGroupAndContent(item.Content)
		if len(todo.Groups[group]) == 0 {
			todo.Groups[group] = []model.Item{}
		}
		todo.Groups[group] = append(todo.Groups[group], item)
	}

	json.NewEncoder(w).Encode(todo)
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
	err := query.DeleteTodo(rs.Service, chi.URLParam(r, "uuid"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Todo deleted"))
}
