package route

import (
	"../handler"
	"github.com/gorilla/mux"
	"net/http"
)

func InitLengKengUserRoute(r *mux.Router, sub *mux.Router) {
	sub.HandleFunc("/lengkeng/view/{id}", handler.ViewLengKengHandler)
	sub.HandleFunc("/lengkeng/new", handler.AddLengKengHandler)
	sub.HandleFunc("/lengkeng/delete/{id}", handler.DeleteLengKengHandler)
	sub.HandleFunc("/lengkeng/update/{id}", handler.UpdateLengKengHandler)
	sub.HandleFunc("/lengkeng/list", handler.ListLengKengHandler)
	sub.HandleFunc("/lengkeng/ping", handler.PingLengKengHandler)
	http.Handle("/api/lengkeng/", r)
}
