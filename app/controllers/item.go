package controllers

import (
	"net/http"

	"github.com/EwenQuim/todo-app/app/common"
	"github.com/EwenQuim/todo-app/app/model"
	"github.com/EwenQuim/todo-app/app/query"
	"github.com/go-chi/chi/v5"
)

func (rs TodoResources) NewItem(w http.ResponseWriter, r *http.Request) {
	newItemCreated, err := common.RequestBody[model.Item](w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newItem, err := query.NewItem(rs.Service, newItemCreated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.SendJSON(w, newItem, http.StatusCreated)
}

func (rs TodoResources) DeleteItem(w http.ResponseWriter, r *http.Request) {
	err := query.DeleteItem(rs.Service, chi.URLParam(r, "itemid")) // c.Params("itemid"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (rs TodoResources) ChangeItem(w http.ResponseWriter, r *http.Request) {
	err := query.ChangeItem(rs.Service, chi.URLParam(r, "itemid"), r.FormValue("new_content")) // c.Params("itemid"))
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
