package controllers

import (
	"net/http"

	"github.com/EwenQuim/todo-app/database"
	"github.com/go-fuego/fuego"
)

type TodoResources struct {
	database.Service
}

func (rs TodoResources) RegisterRoutes(r *fuego.Server) {
	apiRouter := fuego.Group(r, "/api")
	fuego.UseStd(apiRouter, jsonApi)

	fuego.GetStd(apiRouter, "/todo/new", rs.NewTodo)
	fuego.PostStd(apiRouter, "/todo", rs.NewTodo)
	fuego.GetStd(apiRouter, "/todo", rs.GetAllTodos)
	fuego.GetStd(apiRouter, "/todo/{uuid}", rs.GetTodo)
	fuego.GetStd(apiRouter, "/todo/{uuid}/delete", rs.DeleteTodo)

	fuego.PostStd(apiRouter, "/todo/item", rs.NewItem)

	fuego.DeleteStd(apiRouter, "/todo/item/{itemid}", rs.DeleteItem)
	fuego.GetStd(apiRouter, "/todo/{uuid}/switch/{itemid}", rs.SwitchItem)
	fuego.GetStd(apiRouter, "/todo/{uuid}/change/{itemid}", rs.ChangeItem)

	fuego.GetStd(apiRouter, "/item/{itemid}/switch", rs.SwitchItem)
	fuego.GetStd(apiRouter, "/item/{itemid}/change", rs.ChangeItem)

	fuego.GetStd(apiRouter, "/ping", ping)

}

// jsonApi applies the content-type header to all responses of the api as json
func jsonApi(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"response":"pong")}`))
}
