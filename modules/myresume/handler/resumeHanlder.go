package handler

import (
	"../model"
	"../service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func AddResumeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Serve the resource.
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var resume model.Resume
		_ = decoder.Decode(&resume)
		res := service.AddResume(resume)
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

func UpdateResumeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Serve the resource.
	case "POST":

	case "PUT":
		decoder := json.NewDecoder(r.Body)
		var resume model.Resume
		_ = decoder.Decode(&resume)
		res := service.UpdateResume(resume)
		jData, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	case "DELETE":
		// Remove the record.
	default:
		// Give an error message.
	}
}

func DeleteResumeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Serve the resource.
	case "POST":

	case "PUT":

	case "DELETE":
		vars := mux.Vars(r)
		id := vars["id"]
		res := service.DeleteResume(id)
		jData, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	default:
		// Give an error message.
	}
}

func ViewResumeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		id := vars["id"]
		res := service.GetResumeById(id)
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

func ListResumeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		res := service.GetResumeList()
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
