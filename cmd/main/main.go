package main

import (
	"Bookstore/pkg/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	multiplexer := http.NewServeMux()

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	s := &http.Server{
		Addr:    ":8080",
		Handler: multiplexer,
	}
	s.ListenAndServe()
}
