package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func processJson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	interest := vars["activity"]
	todo := Resp{
		Title: "Test Title",
		Body:  interest,
		Href:  "test",
	}
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}
