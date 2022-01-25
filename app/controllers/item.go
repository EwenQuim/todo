package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/EwenQuim/todo-app/app/model"
	"github.com/EwenQuim/todo-app/app/query"
	"github.com/EwenQuim/todo-app/app/validator"
	"github.com/go-chi/chi/v5"
)

func (rs TodoResources) NewItem(w http.ResponseWriter, r *http.Request) {
	newItem := model.Item{
		Content:  validator.CleanItem(r.URL.Query().Get("content")), // c.Query("content")),
		TodoUUID: chi.URLParam(r, "uuid"),                           // c.Params("uuid"),
	}

	newItem, err := query.NewItem(rs.Service, newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}

func (rs TodoResources) DeleteItem(w http.ResponseWriter, r *http.Request) {
	err := query.DeleteItem(rs.Service, chi.URLParam(r, "itemid")) // c.Params("itemid"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (rs TodoResources) SwitchItem(w http.ResponseWriter, r *http.Request) {
	err := query.SwitchItem(rs.Service, chi.URLParam(r, "itemid")) // c.Params("itemid"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
