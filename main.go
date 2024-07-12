package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"log"
	"github.com/gorilla/mux"
)

type AzLoginOutput struct {
	CloudName          string `json:"cloudName"`
	HomeTenantId       string `json:"homeTenantId"`
	Id                 string `json:"id"`
	IsDefault          bool   `json:"isDefault"`
	ManagedByTenants   []interface{} `json:"managedByTenants"`
	Name               string `json:"name"`
	State              string `json:"state"`
	TenantDefaultDomain string `json:"tenantDefaultDomain"`
	TenantDisplayName  string `json:"tenantDisplayName"`
	TenantId           string `json:"tenantId"`
	User               struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"user"`
}

func executeAzLogin() (string, error) {
	// Execute the az login command
	cmd := exec.Command("az", "login")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error: %s", stderr.String())
	}

	// Parse the JSON result
	var output []AzLoginOutput
	err = json.Unmarshal(out.Bytes(), &output)
	if err != nil {
		return "", fmt.Errorf("error parsing JSON: %s", err)
	}

	// Extract the user name before the @ sign
	if len(output) > 0 {
		email := output[0].User.Name
		parts := strings.Split(email, "@")
		if len(parts) > 0 {
			userName := parts[0]
			return userName, nil
		} else {
			return "", fmt.Errorf("invalid email format")
		}
	} else {
		return "", fmt.Errorf("no login information found")
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	userName, err := executeAzLogin()
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

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/login", loginHandler).Methods("GET")
	
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
