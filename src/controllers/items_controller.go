package controllers

import (
	"net/http"

	"github.com/aditya43/bookstore-oauth-go/oauth"
	"github.com/aditya43/golang-bookstore_items-api/src/domain/items"
	"github.com/aditya43/golang-bookstore_items-api/src/services"
	"github.com/aditya43/golang-bookstore_items-api/src/utils/errors"
	"github.com/aditya43/golang-bookstore_items-api/src/utils/http_utils"
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
		http_utils.SendErrorResponse(w, errors.RESTErr{
			Message: err.Message,
			Status:  err.Status,
			Error:   err.Error,
		})
		return
	}

	item := &items.Item{
		Seller: oauth.GetUserId(r),
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		http_utils.SendErrorResponse(w, *err)
		return
	}

	http_utils.SendJsonResponse(w, http.StatusCreated, result)
}

func (cont *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
