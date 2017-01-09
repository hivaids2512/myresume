package resume

import (
	"../../db"
	"github.com/op/go-logging"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var log = logging.MustGetLogger("myresume")

func AddResume(resume Resume) GeneralResponse {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("myresume").C("resumes")
	err := collection.Insert(resume)
	res := GeneralResponse{}
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

func GetResumeById(id string) GeneralResponse {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("myresume").C("resumes")
	result := Resume{}
	err := collection.Find(bson.M{"_id": id}).One(&result)
	res := GeneralResponse{}
	if err != nil {
		log.Error(err)
		res.Error = err
		res.Success = false
	} else {
		res.Data = result
		res.Success = true
		res.Error = err
	}
	return res
}

func DeleteResume(id string) GeneralResponse {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("myresume").C("resumes")
	err := collection.Remove(bson.M{"_id": id})
	res := GeneralResponse{}
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

func UpdateResume(resume Resume) GeneralResponse {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("myresume").C("resumes")
	err := collection.Update(bson.M{"_id": resume.Id}, bson.M{"$set": resume})
	res := GeneralResponse{}
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

func GetResumeList() GeneralResponse {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("myresume").C("resumes")
	result := []Resume{}
	err := collection.Find(nil).All(&result)
	res := GeneralResponse{}
	if err != nil {
		log.Error(err)
		res.Error = err
		res.Success = false
	} else {
		res.List = result
		res.Success = true
		res.Error = err
	}
	return res
}
