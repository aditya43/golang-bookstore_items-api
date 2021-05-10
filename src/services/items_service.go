package services

import (
	"github.com/aditya43/golang-bookstore_items-api/src/domain/items"
	"github.com/aditya43/golang-bookstore_items-api/src/utils/errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(*items.Item) (*items.Item, *errors.RESTErr)
	Get(string) (*items.Item, *errors.RESTErr)
}

type itemsService struct{}

func (s *itemsService) Create(item *items.Item) (*items.Item, *errors.RESTErr) {
	return nil, nil
}
func (s *itemsService) Get(id string) (*items.Item, *errors.RESTErr) {
	return nil, nil
}
