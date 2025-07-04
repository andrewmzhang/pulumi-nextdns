package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func createRewrite(apiKey string, profileId string, name string, content string) (string, error) {
	url := fmt.Sprintf("https://api.nextdns.io/profiles/%s/rewrites/", profileId)

	// Request payload
	payload := map[string]string{
		"name":    name,
		"content": content,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Add headers
	req.Header.Set("X-Api-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Check for HTTP errors
	if resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		log.Fatalf("HTTP error: %s\n%s", resp.Status, string(body))
	}

	// Parse response JSON
	var result struct {
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}

	return result.Data.ID, nil
}

func deleteRewrite(apiKey string, profileId string, rewriteId string) error {
	url := fmt.Sprintf("https://api.nextdns.io/profiles/%s/rewrites/%s", profileId, rewriteId)

	// Create HTTP request
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Add headers
	req.Header.Set("X-Api-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Check for HTTP errors
	if resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		log.Fatalf("HTTP error: %s\n%s", resp.Status, string(body))
	}
	return nil
}
