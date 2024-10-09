package routes

import (
	"github.com/Bhushan9001/GO_CRUD/internal/controllers"
	"github.com/gorilla/mux"
)

func authRoutes(r *mux.Router) {

	authRouter := r.PathPrefix("/auth").Subrouter();

	authRouter.HandleFunc("/signup",controllers.Signup).Methods("POST");
	authRouter.HandleFunc("/signin",controllers.Signin).Methods("POST");

}