package controllers

import (
	"net/http"

	"github.com/EwenQuim/todo-app/database"
	"github.com/go-chi/chi/v5"
)

type TodoResources struct {
	database.Service
}

func (rs TodoResources) RegisterRoutes(r *chi.Mux) {
	apiRouter := chi.NewRouter()
	apiRouter.Use(jsonApi)

	apiRouter.Get("/todo/new", rs.NewTodo)
	apiRouter.Get("/todo", rs.GetAllTodos)
	apiRouter.Get("/todo/{uuid}", rs.GetTodo)
	apiRouter.Get("/todo/{uuid}/delete", rs.DeleteTodo)

	apiRouter.Get("/todo/{uuid}/new", rs.NewItem)
	apiRouter.Get("/todo/{uuid}/delete/{itemid}", rs.DeleteItem)
	apiRouter.Get("/todo/{uuid}/switch/{itemid}", rs.SwitchItem)
	apiRouter.Get("/todo/{uuid}/change/{itemid}", rs.ChangeItem)

	apiRouter.Get("/item/{itemid}/delete", rs.DeleteItem)
	apiRouter.Get("/item/{itemid}/switch", rs.SwitchItem)
	apiRouter.Get("/item/{itemid}/change", rs.ChangeItem)

	apiRouter.Get("/ping", ping)

	r.Mount("/api", apiRouter)
}

// jsonApi applies the content-type header to all responses of the api as json
func jsonApi(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
