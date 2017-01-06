package user

import (
	"../../components/auth"
	"../../config/global"
	"../../db"
	"github.com/op/go-logging"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var log = logging.MustGetLogger("myresume")

func Register(user User) RegisterResponse {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("myresume").C("users")
	user.Password = auth.Hash(user.Password + global.SALT)
	err := collection.Insert(user)
	res := RegisterResponse{}
	if err != nil {
		log.Error(err)
		res.Error = err
		res.Success = false
	} else {
		user.Password = "*********"
		res.User = user
		res.Error = err
		res.Success = true
	}
	return res
}

func Login(user User) LoginResponse {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("myresume").C("users")
	result := User{}
	err := collection.Find(bson.M{"email": user.Email, "password": auth.Hash(user.Password + global.SALT)}).One(&result)
	res := LoginResponse{}
	if err != nil {
		log.Error(err)
		res.Error = err
	} else {
		token, err1 := auth.GenTokenV2()
		if err1 != nil {
			log.Error(err1)
		}
		result.Password = "*********"
		res.User = result
		res.Token = token
		res.Error = err
	}
	return res
}
