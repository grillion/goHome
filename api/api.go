package api

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
)

type ErrorResponse struct {
	Message string `json:"errorMessage"`
}

func init(){

}

/**
 * Send an error response
 */
func sendServerError(w http.ResponseWriter, errorMessage string){
	w.WriteHeader(http.StatusInternalServerError)
	js, err := json.Marshal(&ErrorResponse{errorMessage})
	if err != nil {
		log.Printf("JSON Error while build error response: %s\n", err.Error())
		w.Write([]byte("{\"errorMessage\": \"Unknown server error occurred.\"}"))
		return
	}
	w.Write(js)
}


/**
 * Send bad request response
 */
func sendBadRequest(w http.ResponseWriter, errorMessage string){
	w.WriteHeader(http.StatusBadRequest)
	js, err := json.Marshal(&ErrorResponse{errorMessage})
	if err != nil {
		log.Printf("JSON Error while build error response: %s\n", err.Error())
		w.Write([]byte("{\"errorMessage\": \"Unknown server error occurred.\"}"))
		return
	}
	w.Write(js)
}

/**
 * Send not authorized
 */
func sendNotAuthorized(w http.ResponseWriter){
	w.WriteHeader(http.StatusUnauthorized)
}

/**
 * Send a successful response with data
 */
func sendData(w http.ResponseWriter, marshaledJson []byte){
	w.WriteHeader(http.StatusOK)
	w.Write(marshaledJson)
}

func sendNoData(w http.ResponseWriter ){
	w.Write([]byte("{\"success\":true}"))
}

/**
 * Public method to install the api routes to the http server
 */
func AddRoutes(r *mux.Router)  {

	// Auth
	r.Handle("/api/auth/login", http.HandlerFunc(handleAuthLogin)).Methods("POST")

	// Get all devices
	r.Handle("/api/devices/", http.HandlerFunc(handleDeviceCreate)).Methods("POST")
	r.Handle("/api/devices/", http.HandlerFunc(handleDeviceGetAll)).Methods("GET")
	r.Handle("/api/devices/", http.HandlerFunc(handleDeviceUpdate)).Methods("PUT")
	r.Handle("/api/devices/{id}", http.HandlerFunc(handleDeviceGet)).Methods("GET")
	r.Handle("/api/devices/{id}", http.HandlerFunc(handleDeviceRemove)).Methods("DELETE")

	// User CRUD operations
	r.Handle("/api/users/", http.HandlerFunc(handleUserCreate)).Methods("POST")
	r.Handle("/api/users/", http.HandlerFunc(handleUserGetAll)).Methods("GET")
	r.Handle("/api/users/", http.HandlerFunc(handleUserUpdate)).Methods("PUT")
	r.Handle("/api/users/{id}", http.HandlerFunc(handleUserGet)).Methods("GET")
	r.Handle("/api/users/{id}", http.HandlerFunc(handleUserRemove)).Methods("DELETE")
}