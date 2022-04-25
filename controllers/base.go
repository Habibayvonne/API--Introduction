package controllers

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/innv8/api-introduction/middleware"
	"github.com/innv8/api-introduction/utils"
)

// Base struct is just a struct that holds shared resources
// db connection, router etc
// the side effect of this is, we will also initialize the app in controllers.
// Note: all controllers are now methods to this struct
// this is to ensure that all controllers have access to the resources in the struct.
// the Base struct is used as a pointer in all methods (so that we only have one base)
// all resources in Base should be pointers so that we still have one of each
type Base struct {
	Router *mux.Router
	DB     *sql.DB
}

func (b *Base) StartRouter() {
	b.Router = mux.NewRouter()

	b.Router.HandleFunc("/", middleware.HeadersMiddleware(b.HomeController))
	b.Router.HandleFunc("/users", middleware.HeadersMiddleware(b.UsersController)).Methods("GET", "POST")

}

func (b *Base) StartAPI() {
	var err error
	port := ":5000"

	// connect to db
	b.DB, err = utils.ConnectToDB()
	if err != nil {
		log.Println("unable to connect to db because", err)
		return
	}

	log.Printf("running on port %s", port)
	err = http.ListenAndServe(port, b.Router)
	if err != nil {
		log.Printf("unable to start API because %s", err.Error())
	}
}
