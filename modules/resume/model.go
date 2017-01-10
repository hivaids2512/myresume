package resume

import (
	"gopkg.in/mgo.v2/bson"
)

type Resume struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	Description string        `bson:"description"`
	UserId      bson.ObjectId `json:"userId" bson:"userId,omitempty"`
}

type GeneralResponse struct {
	Success bool
	Error   error
	Data    Resume
	List    []Resume
}
