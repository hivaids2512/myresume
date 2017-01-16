package handler

import (
	"../model"
	"../service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func AddSectionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Serve the resource.
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var section model.Section
		_ = decoder.Decode(&section)
		res := service.AddSection(section)
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

func UpdateSectionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Serve the resource.
	case "POST":

	case "PUT":
		decoder := json.NewDecoder(r.Body)
		var section model.Section
		_ = decoder.Decode(&section)
		res := service.UpdateSection(section)
		jData, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	case "DELETE":
		// Remove the record.
	default:
		// Give an error message.
	}
}

func DeleteSectionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Serve the resource.
	case "POST":

	case "PUT":

	case "DELETE":
		vars := mux.Vars(r)
		id := vars["sectionId"]
		res := service.DeleteSection(id)
		jData, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	default:
		// Give an error message.
	}
}

func ViewSectionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		id := vars["sectionId"]
		res := service.GetSectionById(id)
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

func ListSectionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		resumeId := vars["resumeId"]
		res := service.GetSectionList(resumeId)
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
