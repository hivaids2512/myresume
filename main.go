package main

import (
	"./modules/lengkeng/route"
	myResumeRoute "./modules/myresume/route"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	sub := r.PathPrefix("/api").Subrouter()
	myResumeRoute.InitResumeRoute(r, sub)
	myResumeRoute.InitUserRoute(r, sub)
	route.InitLengKengUserRoute(r, sub)
	var port string
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	} else {
		port = ":8080"
	}
	http.ListenAndServe(port, nil)
}
