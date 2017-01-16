package service

import (
	"../../../db"
	"../model"
	"../model/response"
	"github.com/op/go-logging"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var log = logging.MustGetLogger("myresume")

func AddResume(resume model.Resume) response.Status {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("myresume").C("resumes")
	err := collection.Insert(resume)
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

func GetResumeById(id string) model.DataResponse {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("myresume").C("resumes")
	result := model.Resume{}
	err := collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	res := model.DataResponse{}
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

func DeleteResume(id string) response.Status {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("myresume").C("resumes")
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

func UpdateResume(resume model.Resume) response.Status {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("myresume").C("resumes")
	err := collection.Update(bson.M{"_id": resume.Id}, bson.M{"$set": resume})
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

func GetResumeList() model.ListResponse {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("myresume").C("resumes")
	result := []model.Resume{}
	err := collection.Find(nil).All(&result)
	res := model.ListResponse{}
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
