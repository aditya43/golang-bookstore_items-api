package controllers

import (
	"encoding/json"
	"io/ioutil"
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

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restErr := errors.BadRequestErr("Invalid request body")
		http_utils.SendErrorResponse(w, *restErr)
		return
	}
	defer r.Body.Close()

	var itemRequest *items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		restErr := errors.BadRequestErr("Invalid JSON body")
		http_utils.SendErrorResponse(w, *restErr)
		return
	}

	itemRequest.Seller = oauth.GetUserId(r)

	result, createErr := services.ItemsService.Create(itemRequest)
	if err != nil {
		http_utils.SendErrorResponse(w, *createErr)
		return
	}

	http_utils.SendJsonResponse(w, http.StatusCreated, result)
}

func (cont *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
