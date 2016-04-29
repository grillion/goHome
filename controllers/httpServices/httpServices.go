package httpServices

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
	"github.com/grillion/goHome/config"
	"github.com/grillion/goHome/api"
)

type ErrorResponse struct {
	Message string `json:"errorMessage"`
}

func init(){

	r := mux.NewRouter()

	api.AddRoutes(r)

	addStaticFiles(r)


	http.Handle("/", r)
}

func addStaticFiles(r *mux.Router){
	// Serve the static dir
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(config.GetWebAppRoot())))
	log.Printf("    Static files added to http services")
}

func Start(){
	http.ListenAndServe(":3000", nil);
	log.Printf("Running Web Server on localhost:3000")
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
