package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type BookingModel struct {
	ID primitive.ObjectID `bson:"_id"`
	Booking_id string `json:"booking_id"`;
	User_name string `json:"user_name"`
	User_phone_number string `json:"user_phone_number"`
	Paid_status string `json:"paid_status"`
	Booked_seat string `json:"booked_seat"`
}