package service

import (
	"../../../components/auth"
	"../../../config/global"
	"../../../db"
	"../model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Register(user model.User) model.RegisterResponse {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("myresume").C("users")
	user.Password = auth.Hash(user.Password + global.SALT)
	err := collection.Insert(user)
	res := model.RegisterResponse{}
	if err != nil {
		log.Error(err)
		res.Status.Error = err
		res.Status.Success = false
	} else {
		user.Password = "*********"
		res.User = user
		res.Status.Error = err
		res.Status.Success = true
	}
	return res
}

func Login(user model.User) model.LoginResponse {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("myresume").C("users")
	result := model.User{}
	err := collection.Find(bson.M{"email": user.Email, "password": auth.Hash(user.Password + global.SALT)}).One(&result)
	res := model.LoginResponse{}
	if err != nil {
		log.Error(err)
		res.Status.Error = err
		res.Status.Success = false
	} else {
		token, err1 := auth.GenTokenV2()
		if err1 != nil {
			log.Error(err1)
		}
		result.Password = "*********"
		res.User = result
		res.Token = token
		res.Status.Error = err
		res.Status.Success = true
	}
	return res
}
