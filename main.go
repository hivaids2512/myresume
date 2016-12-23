package main

import (
	"./modules/user"
	"net/http"
)

func main() {
	user.Gonow()
	http.ListenAndServe(":8080", nil)
}
