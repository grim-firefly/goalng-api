package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/grim-firefly/golang-api/pkg/routes"
)

func main() {
	router := chi.NewRouter()
	routes.RegisterBookRoutes(router)
	log.Fatal(http.ListenAndServe(":3000", router))

}
