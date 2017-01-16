package route

import (
	"../handler"
	"github.com/gorilla/mux"
)

func InitResumeRoute(r *mux.Router) {
	r.HandleFunc("/resumes/view/{id}", handler.ViewResumeHandler)
	r.HandleFunc("/resumes/new", handler.AddResumeHandler)
	r.HandleFunc("/resumes/delete/{id}", handler.DeleteResumeHandler)
	r.HandleFunc("/resumes/update/{id}", handler.UpdateResumeHandler)
	r.HandleFunc("/resumes/list", handler.ListResumeHandler)
}
