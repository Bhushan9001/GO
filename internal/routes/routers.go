package routes

import "github.com/gorilla/mux"

func Routes() *mux.Router {
     
	router := mux.NewRouter();

    authRoutes(router);

	return router;
}