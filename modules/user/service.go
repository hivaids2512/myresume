package user

import (
	"../../db"
	"github.com/op/go-logging"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var log = logging.MustGetLogger("example")

func Register(user User) (User, error) {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("myresume").C("users")
	err := collection.Insert(user)
	if err != nil {
		log.Error(err)
	}
	return user, err
}

func Login(user User) (User, error) {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("myresume").C("users")
	result := User{}
	err := collection.Find(bson.M{"email": user.Email, "password": user.Password}).One(&result)
	if err != nil {
		log.Error(err)
	}
	return result, err
}
