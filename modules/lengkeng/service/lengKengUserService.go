package service

import (
	"../../../db"
	"../../myresume/model/response"
	"../model"
	"github.com/op/go-logging"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var log = logging.MustGetLogger("myresume")

func AddLengKengUser(lengkengUser model.LengKengUser) response.Status {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("myresume").C("lengkengusers")
	newId := bson.NewObjectId()
	lengkengUser.Id = newId
	err := collection.Insert(lengkengUser)
	res := response.Status{}
	if err != nil {
		log.Error(err)
		res.Error = err
		res.Success = false
	} else {
		res.Error = err
		res.Success = true
		res.CustomData = newId
	}
	return res
}

func GetLengKengUserById(id string) model.DataResponse {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("myresume").C("lengkengusers")
	result := model.LengKengUser{}
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

func GetLengKengUserByUserId(id string) model.DataResponse {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("myresume").C("lengkengusers")
	result := model.LengKengUser{}
	err := collection.Find(bson.M{"userId": id}).One(&result)
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

func DeleteLengKengUser(id string) response.Status {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("lengkeng").C("lengkengusers")
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

func UpdateLengKengUser(lengKengUser model.LengKengUser) response.Status {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("lengkeng").C("lengkengusers")
	err := collection.Update(bson.M{"_id": lengKengUser.Id}, bson.M{"$set": lengKengUser})
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

func GetLengKengUserList() model.ListResponse {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("lengkeng").C("lengkengusers")
	result := []model.LengKengUser{}
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
