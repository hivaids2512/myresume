package model

import (
	"../../myresume/model/response"
	"gopkg.in/mgo.v2/bson"
)

type LengKengUser struct {
	Id     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserId string        `bson:"userId"`
	Token  string        `bson:"description"`
}

type DataResponse struct {
	Status response.Status
	Data   LengKengUser
}

type ListResponse struct {
	Status response.Status
	List   []LengKengUser
}
