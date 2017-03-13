package handler

import (
	"../model"
	"../service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func AddLengKengHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Serve the resource.
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var lengkengUser model.LengKengUser
		_ = decoder.Decode(&lengkengUser)
		res := service.AddLengKengUser(lengkengUser)
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

func UpdateLengKengHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Serve the resource.
	case "POST":

	case "PUT":
		decoder := json.NewDecoder(r.Body)
		var lengKengUser model.LengKengUser
		_ = decoder.Decode(&lengKengUser)
		res := service.UpdateLengKengUser(lengKengUser)
		jData, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	case "DELETE":
		// Remove the record.
	default:
		// Give an error message.
	}
}

func DeleteLengKengHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Serve the resource.
	case "POST":

	case "PUT":

	case "DELETE":
		vars := mux.Vars(r)
		id := vars["id"]
		res := service.DeleteLengKengUser(id)
		jData, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	default:
		// Give an error message.
	}
}

func ViewLengKengHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		id := vars["id"]
		res := service.GetLengKengUserById(id)
		jData, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	case "POST":

	case "PUT":

	case "DELETE":

	default:
		// Give an error message.
	}
}

func ListLengKengHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		res := service.GetLengKengUserList()
		jData, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	case "POST":

	case "PUT":

	case "DELETE":

	default:
		// Give an error message.
	}
}
