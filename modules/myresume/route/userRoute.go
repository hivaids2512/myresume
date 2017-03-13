package route

import (
	"../handler"
	"github.com/gorilla/mux"
	"net/http"
)

func InitUserRoute(r *mux.Router, sub *mux.Router) {
	sub.HandleFunc("/users/register", handler.RegisterHandler)
	sub.HandleFunc("/users/login", handler.LoginHandlder)
	http.Handle("/api/users/", r)
}
