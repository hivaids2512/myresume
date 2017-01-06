package main

import (
	"./modules/resume"
	"./modules/user"
	"net/http"
	"os"
)

func main() {
	user.Gonow()
	resume.Gonow()
	var port string
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	} else {
		port = ":8080"
	}
	http.ListenAndServe(port, nil)
}
