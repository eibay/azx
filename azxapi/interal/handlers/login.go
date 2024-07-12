// internal/handlers/login.go
package handlers

import (
	"encoding/json"
	"net/http"
	"azxapi/internal/services"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	userName, err := services.ExecuteAzLogin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"userName": userName}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "error creating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
