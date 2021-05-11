package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/aditya43/bookstore-oauth-go/oauth"
	"github.com/aditya43/golang-bookstore_items-api/src/domain/items"
	"github.com/aditya43/golang-bookstore_items-api/src/services"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct {
}

func (cont *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(err.Status)
		json.NewEncoder(w).Encode(err)
		return
	}

	item := &items.Item{
		Seller: oauth.GetUserId(r),
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(err.Status)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func (cont *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
