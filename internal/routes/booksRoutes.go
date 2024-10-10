package routes

import (
	"github.com/Bhushan9001/GO_CRUD/internal/controllers"
	"github.com/gorilla/mux"
)

func bookRoutes(r *mux.Router){

	bookRouter := r.PathPrefix("/books").Subrouter();

	bookRouter.HandleFunc("/",controllers.GetBooks).Methods("GET")
	bookRouter.HandleFunc("/{id}",controllers.GetBook).Methods("GET")
	bookRouter.HandleFunc("/",controllers.CreateBook).Methods("POST")
	bookRouter.HandleFunc("/{id}",controllers.UpdateBook).Methods("PUT")
	bookRouter.HandleFunc("/{id}",controllers.DeleteBook).Methods("DELETE")
}