package models

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"id"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Email    string             `json:"email,omitempty" validate:"required"`
	Password []byte             `json:"-"`
}

type Product struct {
	Pro_Id      primitive.ObjectID `json:"id" bson:"id"`
	ProductName string             `json:"productname,omitempty" validate:"required"`
	Cost        json.Number        `json:"cost,omitempty" validate:"required"`
	UserMail    string             `json:"mail,omitempty"`
}
