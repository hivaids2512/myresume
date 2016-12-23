package main

import (
	//"./config/db" //local package
	"./db"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id      string
	Balance uint64
}

func handler(w http.ResponseWriter, r *http.Request) {
	a := db.GetMongoSession()
	fmt.Println(a)
	u := User{Id: "US123", Balance: 8}
	js, err := json.Marshal(u)
	fmt.Println(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	var quy string = "quy"
	a := "123"
	fmt.Println(a)
	fmt.Println(quy)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
