package main

import (
	"./modules/user"
	"net/http"
	"os"
)

func main() {
	user.Gonow()
	var port string
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	} else {
		port = ":8080"
	}
	http.ListenAndServe(port, nil)
}
