package model

import (
	"./response"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Email    string        `bson:"email"`
	Password string        `bson:"password"`
}

type LoginResponse struct {
	Status response.Status
	User   User
	Token  string
}

type RegisterResponse struct {
	Status response.Status
	User   User
}
