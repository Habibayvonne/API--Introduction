package controllers

import (
	"log"
	"net/http"

	"github.com/innv8/api-introduction/middleware"
	"github.com/innv8/api-introduction/models"
)

// http.ResponseWriter is used to write the response to the client
// http.Request pointer has all the data from the client
func (b *Base) MembersController(w http.ResponseWriter, r *http.Request) {

	log.Println("the connection that reached the controller ", r.Header.Get("Connection"))

	members, _ := models.FetchMembers(b.DB)
	middleware.JSONResponse(w, 200, members)
}

func (b *Base) MembersByPostionController(w http.ResponseWriter, r *http.Request) {

	log.Println("the connection that reached the controller ", r.Header.Get("Connection"))

	members, _ := models.FetchMembersByPosition(1, b.DB)
	middleware.JSONResponse(w, 200, members)
}
