package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/innv8/api-introduction/models"
)

func UsersController(w http.ResponseWriter, r *http.Request) {
	users, _ := models.GetUsers()

	responseJSON, _ := json.Marshal(users)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(200)
	_, _ = w.Write(responseJSON)
}
