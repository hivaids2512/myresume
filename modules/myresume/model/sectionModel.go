package model

import (
	"./response"
	"gopkg.in/mgo.v2/bson"
)

type Section struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title       string        `bson:"title"`
	Description string        `bson:"description"`
	Content     string        `bson:"content"`
	ResumeId    bson.ObjectId `json:"resumeId" bson:"resumeId,omitempty"`
}

type SectionDataResponse struct {
	Status response.Status
	Data   Section
}

type SectionListResponse struct {
	Status response.Status
	List   []Section
}
