package main

import (
	"./modules/myresume/route"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	route.InitResumeRoute(r)
	route.InitUserRoute(r)
	route.InitSectionRoute(r)
	http.Handle("/", r)
	var port string
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	} else {
		port = ":8080"
	}
	http.ListenAndServe(port, nil)
}
