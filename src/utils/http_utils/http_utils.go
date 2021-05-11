package http_utils

import (
	"encoding/json"
	"net/http"

	"github.com/aditya43/golang-bookstore_items-api/src/utils/errors"
)

func SendJsonResponse(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func SendErrorResponse(w http.ResponseWriter, err errors.RESTErr) {
	SendJsonResponse(w, err.Status, err)
}
