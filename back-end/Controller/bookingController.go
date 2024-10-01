package controller

import (
	"Backend/Model"
	"encoding/json"
	"log"
	"net/http"
	database "Backend/Database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func BookShowByShowId(w http.ResponseWriter, r *http.Request){
	if(r.Method == http.MethodPost){
		w.WriteHeader(http.StatusBadRequest);
		json.NewEncoder(w).Encode(status{"error":"supposed to be POST request"});
		return;
	}

	var booking_model *model.BookingModel;

	err:= json.NewDecoder(r.Body).Decode(&booking_model);

	if err!=nil{
		log.Println("failed to parse the request body")
		w.WriteHeader(http.StatusBadRequest);
		json.NewEncoder(w).Encode(status{"error":"failed to parse the request body"});
		return;
	}

	booking_model.ID = primitive.NewObjectID();
	booking_model.Booking_id = booking_model.ID.Hex();

	if booking_model.Show_id==""{
		w.WriteHeader(http.StatusNotFound);
		json.NewEncoder(w).Encode(status{"error":"the showe-id does not exist"})
		return;
	}

	//checking already registered seat is not required because , the validation will be 
	//added in front end itself.

	successId,err := database.CreateBookingByshowId(booking_model.Show_type,
			booking_model.Show_id,
			booking_model.Booked_seat,booking_model)

	
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError);
		json.NewEncoder(w).Encode(status{"error":"something went wrong"});
		return;
	}

	w.WriteHeader(http.StatusOK);
	json.NewEncoder(w).Encode(status{"message":"success","data":successId});
	return;
}