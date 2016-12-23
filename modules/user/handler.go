package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	u := User{Id: "US123", Balance: 8}
	js, err := json.Marshal(u)
	fmt.Println(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
