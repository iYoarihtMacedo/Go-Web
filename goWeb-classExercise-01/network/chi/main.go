package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	// router
	router := chi.NewRouter()

	// endpoints
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		// request
		fmt.Println("GET /ping")
		fmt.Println("method:", r.Method)
		fmt.Println("url:", r.URL)
		fmt.Println("header:", r.Header)

		// response
		w.Write([]byte(`{"message":"pong"}`))
	})

	router.Post("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message":"pong", "method":"POST"}`))
	})

	http.ListenAndServe(":8080", router)
}
