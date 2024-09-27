package routes

import (
	// "net/http"
	"github.com/gorilla/mux"
	"Backend/Controller"
)

func AdminRouter() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("create-movie",controller.CreateShowController).Methods("POST");

	return router;
}