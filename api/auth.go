package api

import (
	"net/http"
	"encoding/json"
)

type LoginSubmission struct {
	username, password string
}


func handleAuthLogin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var newUser LoginSubmission
	err := decoder.Decode(&newUser)

	if err != nil {
		sendServerError(w, err.Error())
		return
	}

}
