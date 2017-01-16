package route

import (
	"../handler"
	"github.com/gorilla/mux"
)

func InitSectionRoute(r *mux.Router) {
	r.HandleFunc("/resumes/{resumeId}/sections/view/{sectionId}", handler.ViewSectionHandler)
	r.HandleFunc("/resumes/{resumeId}/sections/new", handler.AddSectionHandler)
	r.HandleFunc("/resumes/{resumeId}/sections/delete/{sectionId}", handler.DeleteSectionHandler)
	r.HandleFunc("/resumes/{resumeId}/sections/update/{sectionId}", handler.UpdateSectionHandler)
	r.HandleFunc("/resumes/{resumeId}/sections/list", handler.ListSectionHandler)
}
