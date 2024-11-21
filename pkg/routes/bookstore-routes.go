package routes

import (
	"github.com/dimnyan/bookstore-api/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func (router *mux.Router){
	router.HandleFunc("/api/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/api/books", controllers.GetBook).Methods("GET")
	router.HandleFunc("/api/books/{id}", controllers.GetBookById).Methods("GET")
	// router.HandleFunc("/api/books/{id}", controllers.UpdateBooks).Methods("PUT")
	// router.HandleFunc("/api/books/{id}", controllers.DeleteBook).Methods("DELETE")
}