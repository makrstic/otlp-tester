package http

import (
	"bytes"
	"fmt"
	"net/http"
	"otlp-tester/models"
)

func PostRequest(url string, payload bytes.Buffer, auth models.Auth) (*http.Response, error) {
	req, _ := http.NewRequest("POST", url, &payload)

	if auth.AuthType != "none" {
		req.Header.Set("Authorization", auth.AuthType+" "+auth.Token)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Failed to create API call: %w", err)
	}

	// Caller must handle res.Body.Close()
	return res, nil
}
