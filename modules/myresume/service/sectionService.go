package service

import (
	"../../../db"
	"../model"
	"../model/response"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func AddSection(section model.Section) response.Status {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("myresume").C("sections")
	err := collection.Insert(section)
	res := response.Status{}
	if err != nil {
		log.Error(err)
		res.Error = err
		res.Success = false
	} else {
		res.Error = err
		res.Success = true
	}
	return res
}

func GetSectionById(id string) model.SectionDataResponse {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("myresume").C("sections")
	result := model.Section{}
	err := collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	res := model.SectionDataResponse{}
	if err != nil {
		log.Error(err)
		res.Status.Error = err
		res.Status.Success = false
	} else {
		res.Data = result
		res.Status.Success = true
		res.Status.Error = err
	}
	return res
}

func DeleteSection(id string) response.Status {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("myresume").C("sections")
	err := collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	res := response.Status{}
	if err != nil {
		log.Error(err)
		res.Error = err
		res.Success = false
	} else {
		res.Success = true
		res.Error = err
	}
	return res
}

func UpdateSection(section model.Section) response.Status {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("myresume").C("sections")
	err := collection.Update(bson.M{"_id": section.Id}, bson.M{"$set": section})
	res := response.Status{}
	if err != nil {
		log.Error(err)
		res.Error = err
		res.Success = false
	} else {
		res.Success = true
		res.Error = err
	}
	return res
}

func GetSectionList(resumeId string) model.SectionListResponse {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("myresume").C("sections")
	result := []model.Section{}
	err := collection.Find(bson.M{"resumeId": bson.ObjectIdHex(resumeId)}).All(&result)
	res := model.SectionListResponse{}
	if err != nil {
		log.Error(err)
		res.Status.Error = err
		res.Status.Success = false
	} else {
		res.List = result
		res.Status.Success = true
		res.Status.Error = err
	}
	return res
}
