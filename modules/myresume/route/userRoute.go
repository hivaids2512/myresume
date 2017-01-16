package route

import (
	"../handler"
	"github.com/gorilla/mux"
)

func InitUserRoute(r *mux.Router) {
	r.HandleFunc("/users/register", handler.RegisterHandler)
	r.HandleFunc("/users/login", handler.LoginHandlder)
}
