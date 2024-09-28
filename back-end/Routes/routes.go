package routes

import (
	// "net/http"
	"Backend/Controller"
	ratelimiter "Backend/RateLimiter"

	"github.com/gorilla/mux"
)

func AdminRouter() *mux.Router{
	router := mux.NewRouter()

	//with rate limiter
	router.Handle("/create-movie",ratelimiter.RateLimiter(controller.CreateShowController)).Methods("POST");

	return router;
}