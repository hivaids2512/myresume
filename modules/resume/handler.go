package resume

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func addResumeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Serve the resource.
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var resume Resume
		_ = decoder.Decode(&resume)
		res := AddResume(resume)
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

func updateResumeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Serve the resource.
	case "POST":

	case "PUT":
		decoder := json.NewDecoder(r.Body)
		var resume Resume
		_ = decoder.Decode(&resume)
		res := UpdateResume(resume)
		jData, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	case "DELETE":
		// Remove the record.
	default:
		// Give an error message.
	}
}

func deleteResumeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Serve the resource.
	case "POST":

	case "PUT":

	case "DELETE":
		vars := mux.Vars(r)
		id := vars["id"]
		res := DeleteResume(id)
		jData, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	default:
		// Give an error message.
	}
}

func viewResumeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		id := vars["id"]
		res := GetResumeById(id)
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

func listResumeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		res := GetResumeList()
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
