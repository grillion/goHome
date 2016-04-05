package db

import (
	"gopkg.in/mgo.v2"
	"github.com/grillion/goHome/config"
)

var (
	mongoSession *mgo.Session
)

func init(){
	session, err := mgo.Dial(config.GetMongoHost())

	if err != nil {
		panic("Cannot connect to DB")
	}

	mongoSession = session
}

func CloseSession(){
	mongoSession.Close()
}

func getC(collection string) *mgo.Collection {
	return mongoSession.Clone().DB(config.GetMongoDBName()).C(collection)
}