package main

import (
	"log"
	"net/http"

	"github.com/jeki-aka-zer0/go-crud/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
