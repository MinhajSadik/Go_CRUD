package Models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Student struct {
	ID primitive.ObjectID 	`json:"_id,omitempty" bson:"_id,omitempty"`
	Name 		string 		`json:"name"`
	Age 		int 		`json:"age"`
	PhoneNumber string 		`json:"phonenumber"`
}