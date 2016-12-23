package db

import "gopkg.in/mgo.v2"
import "log"

var mgoSession *mgo.Session

// Creates a new session if mgoSession is nil i.e there is no active mongo session.
//If there is an active mongo session it will return a Clone
func GetMongoSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial("mongodb://quyquy2512:123456@ds023438.mlab.com:23438/myresume")
		if err != nil {
			log.Fatal("Failed to start the Mongo session")
		}
	}
	return mgoSession.Clone()
}
