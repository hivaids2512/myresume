package resume

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Gonow() {
	r := mux.NewRouter()
	r.HandleFunc("/resumes/view/{id}", viewResumeHandler)
	r.HandleFunc("/resumes/new", addResumeHandler)
	r.HandleFunc("/resumes/delete/{id}", deleteResumeHandler)
	r.HandleFunc("/resumes/update/{id}", updateResumeHandler)
	r.HandleFunc("/resumes/list", listResumeHandler)
	http.Handle("/resumes/", r)
}
