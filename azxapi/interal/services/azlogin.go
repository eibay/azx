// internal/services/azlogin.go
package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"azxapi/pkg/models"
)

func ExecuteAzLogin() (string, error) {
	cmd := exec.Command("az", "login")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error: %s", stderr.String())
	}

	var output []models.AzLoginOutput
	err = json.Unmarshal(out.Bytes(), &output)
	if err != nil {
		return "", fmt.Errorf("error parsing JSON: %s", err)
	}

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
