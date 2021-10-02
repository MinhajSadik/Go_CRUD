package Models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID primitive.ObjectID 	`json:"_id,omitempty" bson:"_id,omitempty"`
	Name 		string 		`json:"name"`
	Age 		int 		`json:"age"`
	PhoneNumber string 		`json:"phonenumber"`
}


func (obj Student) Validate() error {
	return validation.ValidateStruct(&obj,
	validation.Field(&obj.Name, validation.Required),
	validation.Field(&obj.Age, validation.Required, validation.Min(16), validation.Max(40)),
)
}