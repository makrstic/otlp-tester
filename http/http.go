package http

import (
	"net/http"
	"bytes"
	"fmt"
)

func PostRequest(url string, payload bytes.Buffer) (*http.Response, error) {
	req, _ := http.NewRequest("POST", url, &payload)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Failed to create API call: %w", err)
        }

	// Caller must handle res.Body.Close()
        return res, nil
}
