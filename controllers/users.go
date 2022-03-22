package controllers

import (
	"net/http"

	"github.com/innv8/api-introduction/middleware"
	"github.com/innv8/api-introduction/models"
)

// http.ResponseWriter is used to write the response to the client
// http.Request pointer has all the data from the client
func (b *Base) UsersController(w http.ResponseWriter, r *http.Request) {
	users, _ := models.GetUsers()
	middleware.JSONResponse(w, 200, users)
}
