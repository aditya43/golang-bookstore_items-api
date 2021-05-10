package controllers

import (
	"net/http"

	"github.com/aditya43/bookstore-oauth-go/oauth"
	"github.com/aditya43/golang-bookstore_items-api/src/domain/items"
)

func Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		return
	}

	item := &items.Item{
		Seller: oauth.GetUserId(r),
	}
}

func Get(w http.ResponseWriter, r *http.Request) {

}
