// cmd/server/main.go
package main

import (
	"log"
	"net/http"
	"azxapi/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/login", handlers.LoginHandler).Methods("GET")
	
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
