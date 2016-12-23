package user

import (
	"net/http"
)

func Gonow() {
	http.HandleFunc("/", handler)
}
