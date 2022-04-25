package middleware

import (
	"log"
	"net/http"
)

func HeadersMiddleware(originalController http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("connection header from client", r.Header.Get("Connection"))
		r.Header.Set("connection", "some connection")

		originalController(w, r)

	}
}
