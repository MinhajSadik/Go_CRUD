package Models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name"`
	Age         int                `json:"age"`
	PhoneNumber string             `json:"phonenumber"`
}

func (obj Student) Validate() error {
	return validation.ValidateStruct(&obj,
		validation.Field(&obj.Name, validation.Required),
		validation.Field(&obj.Age, validation.Required, validation.Min(16), validation.Max(40)),
	)
}

func (obj Student) GetBSONModifictionObj() bson.M {
	self := bson.M{
		"name":        obj.Name,
		"age":         obj.Age,
		"phonenumber": obj.PhoneNumber,
	}
	return self
}

type StudentSearch struct {
	Name           string `json:"name"`
	NameIsUsed     bool   `json:"nameisused"`
	AgeMin         int    `json:"agemin"`
	AgeMax         int    `json:"agemax"`
	AgeRangeIsUsed bool   `json:"agerangeisused"`
}

func (obj StudentSearch) GetBSONSearchObj() bson.M {
	self := bson.M{}
	if obj.NameIsUsed {
		self["name"] = obj.Name
	}
	if obj.AgeRangeIsUsed {
		self["age"] = bson.M{
			"$gte": obj.AgeMin,
			"$lte": obj.AgeMax,
		}
	}
	return self
}

type StudentModify struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name"`
	Age         int                `json:"age"`
	PhoneNumber string             `json:"phonenumber"`
}
