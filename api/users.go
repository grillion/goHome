package api

import (
	"net/http"
	"github.com/grillion/goHome/db"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
)

func handleUserCreate(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var newUser db.User
	err := decoder.Decode(&newUser)

	if err != nil {
		sendServerError(w, err.Error())
		return
	}

	// Create User
	Users := db.Users{}
	createdUser, createError := Users.Create(newUser)

	if createError != nil {
		sendServerError(w, createError.Error())
		return
	}

	userJson, jsonErr := json.Marshal(createdUser)
	if jsonErr != nil {
		sendServerError(w, "JSON Error while building response")
		return
	}

	sendData(w, userJson)
}

func handleUserGetAll(w http.ResponseWriter, r *http.Request) {

	Users := db.Users{}
	userList, err := Users.GetAll()

	if err != nil {
		sendServerError(w, err.Error())
		return
	}

	json, jsonErr := json.Marshal(userList)
	if jsonErr != nil {
		sendServerError(w, "JSON Error while building response")
		return
	}

	sendData(w, json)
}

func handleUserGet(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userId := vars["id"]

	if !bson.IsObjectIdHex(userId) {
		sendBadRequest(w, "Invalid User ID")
		return
	}

	bsonId := bson.ObjectIdHex(userId)
	Users := db.Users{}
	userList, err := Users.Get(bsonId)

	if err != nil {
		sendServerError(w, err.Error())
		return
	}

	json, jsonErr := json.Marshal(userList)
	if jsonErr != nil {
		sendServerError(w, "JSON Error while building response")
		return
	}

	sendData(w, json)
}

func handleUserUpdate(w http.ResponseWriter, r *http.Request) {

	Users := db.Users{}

	var updatedUser db.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updatedUser)

	if err != nil {
		sendServerError(w, err.Error())
		return
	}

	updateError := Users.Update(&updatedUser)

	if updateError != nil {
		sendServerError(w, updateError.Error())
		return
	}

	userJson, jsonErr := json.Marshal(updatedUser)
	if jsonErr != nil {
		sendServerError(w, "JSON Error while building response")
		return
	}

	sendData(w, userJson)
}

func handleUserRemove(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userId := vars["id"]

	if !bson.IsObjectIdHex(userId) {
		sendBadRequest(w, "Invalid User ID")
		return
	}

	bsonId := bson.ObjectIdHex(userId)
	Users := db.Users{}
	err := Users.RemoveUser(bsonId)

	if err != nil {
		sendServerError(w, err.Error())
		return
	}

	sendNoData(w);
}