package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Serve the resource.
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var user User
		err := decoder.Decode(&user)
		if err != nil {
			log.Error(err)
		}
		result, err := Register(user)
		log.Info(result.Email)
		if err != nil {
			log.Error(err)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "Success!")
	case "PUT":
		// Update an existing record.
	case "DELETE":
		// Remove the record.
	default:
		// Give an error message.
	}

}

func loginHandlder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Serve the resource.
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var user User
		err := decoder.Decode(&user)
		if err != nil {
			log.Error(err)
		}
		result, err := Login(user)
		log.Info(result.Email)
		if err != nil {
			log.Error(err)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "Success!")
	case "PUT":
		// Update an existing record.
	case "DELETE":
		// Remove the record.
	default:
		// Give an error message.
	}
}
