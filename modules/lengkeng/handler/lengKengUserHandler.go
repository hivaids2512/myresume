package handler

import (
	"../model"
	"../service"
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/op/go-logging"
	"net/http"
)

var log = logging.MustGetLogger("lengkengUser")

func AddLengKengHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Serve the resource.
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var lengkengUser model.LengKengUser
		_ = decoder.Decode(&lengkengUser)
		res := service.AddLengKengUser(lengkengUser)
		jData, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	case "PUT":
		// Update an existing record.
	case "DELETE":
		// Remove the record.
	default:
		// Give an error message.
	}
}

func UpdateLengKengHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Serve the resource.
	case "POST":

	case "PUT":
		decoder := json.NewDecoder(r.Body)
		var lengKengUser model.LengKengUser
		_ = decoder.Decode(&lengKengUser)
		res := service.UpdateLengKengUser(lengKengUser)
		jData, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	case "DELETE":
		// Remove the record.
	default:
		// Give an error message.
	}
}

func DeleteLengKengHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Serve the resource.
	case "POST":

	case "PUT":

	case "DELETE":
		vars := mux.Vars(r)
		id := vars["id"]
		res := service.DeleteLengKengUser(id)
		jData, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	default:
		// Give an error message.
	}
}

func ViewLengKengHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		id := vars["id"]
		res := service.GetLengKengUserById(id)
		jData, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	case "POST":

	case "PUT":

	case "DELETE":

	default:
		// Give an error message.
	}
}

func ListLengKengHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		res := service.GetLengKengUserList()
		jData, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	case "POST":

	case "PUT":

	case "DELETE":

	default:
		// Give an error message.
	}
}

func PingLengKengHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

	case "POST":
		decoder := json.NewDecoder(r.Body)
		var lengkengUserId model.LengKengPing
		_ = decoder.Decode(&lengkengUserId)
		res := service.GetLengKengUserByUserId(lengkengUserId.Id)
		if res.Status.Success == true {
			url := "https://fcm.googleapis.com/fcm/send"
			var jsonStr = []byte(`{ "notification" : {
      "body" : "ahihi!",
      "title" : "Portugal vs. Denmark"
  },
  "to" : "` + res.Data.Token + `"
}`)
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "key=AAAADiQjZEs:APA91bGAmm8MPAfv3xBx9AGLWPUyFB3_2RuWN0UTgRHGB8dvWxybyGy-GzBUgyYu338KUjvRbHNwi7P9j7uqtPyuQTfEVQYDJ_TtsY6_d6jcyQWzdfb1g5dFdR0Y3fXuRhi6QpM8qDwV")
			client := &http.Client{}
			resp, err := client.Do(req)
			log.Error(resp)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
		}

		jData, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	case "PUT":

	case "DELETE":

	default:
		// Give an error message.
	}
}
