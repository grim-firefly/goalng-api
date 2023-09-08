package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/grim-firefly/golang-api/pkg/controllers"
)

// book  routes
var RegisterBookRoutes = func(router *chi.Mux) {
	router.Get("/books", controllers.GetBooks)
	router.Post("/books", controllers.CreateBook)
	router.Get("/books/{id}", controllers.GetBookById)
	router.Put("/books/{id}", controllers.UpdateBook)
	router.Delete("/books/{id}", controllers.DeleteBook)
}
