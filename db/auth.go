package db

import (
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type UserAuth struct {
	UserID bson.ObjectId `bson:"userId"`
	SessionID bson.ObjectId `bson:"sessionId"`
	Created time.Time
}

type Auth struct{}

func (A Auth) Login(username string, password string) (sessionId *bson.ObjectId) {
	Users := Users{}
	user, err := Users.FindOne(bson.M{"username": username, "password": password})

	if err != nil {
		log.Printf("User login error: %s\n", err.Error())
		return nil;
	}

	// Create a new auth object
	newAuth := UserAuth{user.ID, bson.NewObjectId(), time.Now()}

	// Found a valid user, update session table
	getC("session").Upsert(bson.M{"userId": user.ID}, newAuth)

	return
}