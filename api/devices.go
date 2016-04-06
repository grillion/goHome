package api

import (
	"net/http"
	"github.com/grillion/goHome/db"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
)

func handleDeviceCreate(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var newDevice db.Device
	err := decoder.Decode(&newDevice)

	if err != nil {
		sendServerError(w, err.Error())
		return
	}

	// Create Device
	Devices := db.Devices{}
	createdDevice, createError := Devices.Create(newDevice)

	if createError != nil {
		sendServerError(w, createError.Error())
		return
	}

	deviceJson, jsonErr := json.Marshal(createdDevice)
	if jsonErr != nil {
		sendServerError(w, "JSON Error while building response")
		return
	}

	sendData(w, deviceJson)
}

func handleDeviceGetAll(w http.ResponseWriter, r *http.Request) {

	Devices := db.Devices{}
	deviceList, err := Devices.GetAll()

	if err != nil {
		sendServerError(w, err.Error())
		return
	}

	json, jsonErr := json.Marshal(deviceList)
	if jsonErr != nil {
		sendServerError(w, "JSON Error while building response")
		return
	}

	sendData(w, json)
}

func handleDeviceGet(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	deviceId := vars["id"]

	if !bson.IsObjectIdHex(deviceId) {
		sendBadRequest(w, "Invalid Device ID")
		return
	}

	bsonId := bson.ObjectIdHex(deviceId)
	Devices := db.Devices{}
	deviceList, err := Devices.Get(bsonId)

	if err != nil {
		sendServerError(w, err.Error())
		return
	}

	json, jsonErr := json.Marshal(deviceList)
	if jsonErr != nil {
		sendServerError(w, "JSON Error while building response")
		return
	}

	sendData(w, json)
}

func handleDeviceUpdate(w http.ResponseWriter, r *http.Request) {

	Devices := db.Devices{}

	var updatedDevice db.Device
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updatedDevice)

	if err != nil {
		sendServerError(w, err.Error())
		return
	}

	updateError := Devices.Update(&updatedDevice)

	if updateError != nil {
		sendServerError(w, updateError.Error())
		return
	}

	deviceJson, jsonErr := json.Marshal(updatedDevice)
	if jsonErr != nil {
		sendServerError(w, "JSON Error while building response")
		return
	}

	sendData(w, deviceJson)
}

func handleDeviceRemove(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	deviceId := vars["id"]

	if !bson.IsObjectIdHex(deviceId) {
		sendBadRequest(w, "Invalid Device ID")
		return
	}

	bsonId := bson.ObjectIdHex(deviceId)
	Devices := db.Devices{}
	err := Devices.RemoveDevice(bsonId)

	if err != nil {
		sendServerError(w, err.Error())
		return
	}

	sendNoData(w);
}