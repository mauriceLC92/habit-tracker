package main

import (
	"encoding/json"
	"net/http"
)

type Greeting struct {
	Message string `json:"message"`
}

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {

	greeting := Greeting{
		Message: "Hello Havyn team 👋",
	}

	json, err := json.Marshal(greeting)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
