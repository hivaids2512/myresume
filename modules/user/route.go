package user

import (
	"../../components/auth"
	"net/http"
)

func Gonow() {
	http.HandleFunc("/users/register", registerHandler)
	http.HandleFunc("/users/login", loginHandlder)
}
