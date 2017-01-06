package user

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Gonow() {
	r := mux.NewRouter()
	r.HandleFunc("/users/register", registerHandler)
	r.HandleFunc("/users/login", loginHandlder)
	http.Handle("/users/", r)
}
