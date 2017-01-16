package model

import (
	"./response"
	"gopkg.in/mgo.v2/bson"
)

type Resume struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	Description string        `bson:"description"`
	UserId      bson.ObjectId `json:"userId" bson:"userId,omitempty"`
}

type DataResponse struct {
	Status response.Status
	Data   Resume
}

type ListResponse struct {
	Status response.Status
	List   []Resume
}
