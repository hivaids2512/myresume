package user

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Email    string        `bson:"email"`
	Password string        `bson:"password"`
}

type LoginResponse struct {
	User  User
	Token string
	Error error
}

type RegisterResponse struct {
	User    User
	Success bool
	Error   error
}
