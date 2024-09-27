package controller

import "net/http"

func CreateShowController() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		//take values from form
		//get the images urls over gridfs
		//parse it to golang struct
		//store in db
		
	}
}