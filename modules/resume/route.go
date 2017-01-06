package resume

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Gonow() {
	r := mux.NewRouter()
	r.HandleFunc("/resumes/view/{id}", nil)
	r.HandleFunc("/resumes/new", nil)
	r.HandleFunc("/resumes/delete/{id}", nil)
	r.HandleFunc("/resumes/update/{id}", nil)
	r.HandleFunc("/resumes/list", nil)
	http.Handle("/resumes/", r)
}
