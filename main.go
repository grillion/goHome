package main

import (
	"net/http"
	"log"

	"github.com/grillion/goHome/api"
	"github.com/grillion/goHome/config"
	"github.com/gorilla/mux"
	"github.com/grillion/goHome/db"
)

func init() {

}

func main() {

	// Close DB when exiting
	defer db.CloseSession()

	r := mux.NewRouter()

	// Install API Routes to MUX
	api.AddRoutes(r)

	// Serve the static dir
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(config.GetWebAppRoot())))

	log.Printf("Running Web Server on localhost:3000")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
