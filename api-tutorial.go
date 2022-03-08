package main

import (
	"log"
	"net/http"

	"github.com/innv8/api-introduction/controllers"
)

/*
To use 'air' in any project, first run this to make your project a go module

go mod init


If you want to update the packages, you run
go mod tidy
*/

func main() {
	port := ":5000"
	log.Printf("starting API on port %s", port)

	http.HandleFunc("/", controllers.HomeController)
	http.HandleFunc("/users", controllers.UsersController)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Printf("unable to start API because %s", err)
	}
}
