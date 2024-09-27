package main

import (
	 "Backend/Routes"
	// "Backend/Model/Showe/Factory"
	// "Backend/Util"
	// "Backend/Model/Showe"
	"log"
	"net/http"
)

func main(){

	router := routes.AdminRouter()

	log.Println("Starting server on :8080")
	
	if err:=http.ListenAndServe(":8080",router);err!=nil{
		log.Fatal(err)
	}

}