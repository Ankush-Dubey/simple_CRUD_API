package main

import (
	"encoding/json"
	"net/http"
)

func creatbook(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&newBook)
	// newBook.ID = strconv.Itoa(rand.Intn(100))
	// book = append(book, newBook)
	json.NewEncoder(w).Encode(newBook)
}
