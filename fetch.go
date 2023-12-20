package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// HomeHandler is the handler function for the root path ("/")
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the home page!")
}

// HelloHandler is the handler function for the "/hello/{name}" path
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Get the "name" variable from the path
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Println(name)
	res := make(map[string]interface{})
	res["my_Name"] = name
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(res)
}
