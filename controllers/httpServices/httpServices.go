package httpServices

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
	"github.com/grillion/goHome/config"
	"github.com/grillion/goHome/api"
	"fmt"
	"github.com/grillion/mFi"
)

type ErrorResponse struct {
	Message string `json:"errorMessage"`
}

func init(){

}

func addStaticFiles(r *mux.Router){
	// Serve the static dir
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(config.GetWebAppRoot())))
	log.Printf("    Static files added to http services")
}

func Start(){

	appServer := http.NewServeMux()
	appServerRoutes := mux.NewRouter()
	api.AddRoutes(appServerRoutes)
	addStaticFiles(appServerRoutes)
	appServer.Handle("/", appServerRoutes)

	informServer := http.NewServeMux()
	informServerRoutes := mux.NewRouter()
	informServerRoutes.HandleFunc("/inform", func(res http.ResponseWriter , req *http.Request){

		iPkt, err := mFi.ParseInformPacket(req.Body)
		if(nil != err){
			fmt.Printf("Parse Errpr: %s\n", err)
			return
		}
		if(nil == iPkt){
			fmt.Println("Could not parse paylod")
			return
		}

		mJSON, err := json.Marshal(iPkt)
		log.Printf("New Inform Packet: \n%s\n", mJSON)

		//inform.Save(bodyBytes)
	})
	informServer.Handle("/", informServerRoutes)

	// API AND UI
	go func(){ http.ListenAndServe(":3000", appServer); }()
	log.Printf("Running App Server on localhost:3000")

	go func(){ http.ListenAndServe(":6080", informServer); }()
	log.Printf("Running Inform handler localhost:6080")
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
