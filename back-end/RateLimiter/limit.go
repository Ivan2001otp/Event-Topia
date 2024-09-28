package ratelimiter

import (
	"encoding/json"
	"net/http"

	rate "golang.org/x/time/rate"
)

type Message struct{
	status string `json:"status"`
	body   string 	`json:"body"`
}

func RateLimiter(next func(w http.ResponseWriter,r *http.Request))http.Handler{
	limiter := rate.NewLimiter(2,4)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow(){
			message := Message{
				status:"request failed",
				body:"The api is at capacity, try again later.",
			}
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(&message)
			return;	
		}else{
			next(w,r);
		}
	})
}