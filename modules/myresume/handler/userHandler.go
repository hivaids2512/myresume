package handler

import (
	"../model"
	"../service"
	"encoding/json"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Serve the resource.
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var user model.User
		_ = decoder.Decode(&user)
		res := service.Register(user)
		jData, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	case "PUT":
		// Update an existing record.
	case "DELETE":
		// Remove the record.
	default:
		// Give an error message.
	}
}

func LoginHandlder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Serve the resource.
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var user model.User
		_ = decoder.Decode(&user)
		res := service.Login(user)
		jData, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	case "PUT":
		// Update an existing record.
	case "DELETE":
		// Remove the record.
	default:
		// Give an error message.
	}
}
