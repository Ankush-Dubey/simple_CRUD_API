package main

import "github.com/gorilla/mux"

func V1routes(router *mux.Router) {
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/books/list", creatbook).Methods("POST")
}
