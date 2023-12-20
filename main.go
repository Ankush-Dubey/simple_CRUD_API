package main

import (
	"fmt"
	"main/connection"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// server creation
	router := mux.NewRouter()
	// db connection
	connection.Connection()
	// init routes
	V1routes(router)
	// Start the server on port 8080
	port := 8080
	fmt.Printf("Server is listening on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
