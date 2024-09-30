package routes

import (
	// "net/http"
	"Backend/Controller"
	ratelimiter "Backend/RateLimiter"

	"github.com/gorilla/mux"
)

func AdminRouter() *mux.Router{
	router := mux.NewRouter()

	//POST routes
	router.Handle("/create-movie",ratelimiter.RateLimiter(controller.CreateMovieController)).Methods("POST");
	router.Handle("/create-liveshow",ratelimiter.RateLimiter(controller.CreateLiveshowController)).Methods("POST");
	router.Handle("/create-activity",ratelimiter.RateLimiter(controller.CreateActivityController)).Methods("POST");
	router.Handle("/create-event",ratelimiter.RateLimiter(controller.CreateEventController)).Methods("POST");

	//GET routes
	router.Handle("/fetch-movies",ratelimiter.RateLimiter(controller.FetchAllMoviesShowe)).Methods("GET");
	router.Handle("/fetch-activities",ratelimiter.RateLimiter(controller.FetchAllActivity)).Methods("GET");
	router.Handle("/fetch-liveshows",ratelimiter.RateLimiter(controller.FetchAllLiveshow)).Methods("GET");
	router.Handle("/fetch-events",ratelimiter.RateLimiter(controller.FetchAllEvent)).Methods("GET");

	return router;
}