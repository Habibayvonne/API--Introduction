package main

import "github.com/innv8/api-introduction/controllers"

/*
To use 'air' in any project, first run this to make your project a go module

go mod init


If you want to update the packages, you run
go mod tidy
*/

func main() {

	/*
		port := ":5000"
		log.Printf("starting API on port %s", port)

		// we create a router variable and add the routes the same way we did with http
		// then we put router as the second parameter in http.ListenAndServe method
		var router = mux.NewRouter()


		// http.HandleFunc("/", controllers.HomeController)
		// http.HandleFunc("/users", controllers.UsersController)

		err := http.ListenAndServe(port, router)
		if err != nil {
			log.Printf("unable to start API because %s", err)
		}
	*/

	// we are now using a controllers.Base struct to hold the router, this means
	// we also need to start the API from controllers.
	// because if you import a struct from another package, you cannot modify its pointer

	var base controllers.Base
	base.StartRouter()
	base.StartAPI()
}
