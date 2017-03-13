package route

import (
	"../../../components/auth"
	"../handler"
	"github.com/gorilla/mux"
	"net/http"
)

func InitResumeRoute(r *mux.Router, sub *mux.Router) {
	sub.HandleFunc("/resumes/view/{id}", handler.ViewResumeHandler)
	sub.HandleFunc("/resumes/new", handler.AddResumeHandler)
	sub.HandleFunc("/resumes/delete/{id}", handler.DeleteResumeHandler)
	sub.HandleFunc("/resumes/update/{id}", handler.UpdateResumeHandler)
	sub.HandleFunc("/resumes/list", handler.ListResumeHandler)
	sub.HandleFunc("/resumes/{resumeId}/sections/view/{sectionId}", handler.ViewSectionHandler)
	sub.HandleFunc("/resumes/{resumeId}/sections/new", handler.AddSectionHandler)
	sub.HandleFunc("/resumes/{resumeId}/sections/delete/{sectionId}", handler.DeleteSectionHandler)
	sub.HandleFunc("/resumes/{resumeId}/sections/update/{sectionId}", handler.UpdateSectionHandler)
	sub.HandleFunc("/resumes/{resumeId}/sections/list", handler.ListSectionHandler)
	http.Handle("/api/resumes/", auth.ShowMeYouToken(r))
}
