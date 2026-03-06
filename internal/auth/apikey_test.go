package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Test case 1: No Authorization header
	headers := http.Header{}
	_, err := GetAPIKey(headers)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Test case 2: Malformed Authorization header
	headers.Set("Authorization", "Bearer ")
	_, err = GetAPIKey(headers)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Test case 3: Valid Authorization header
	headers.Set("Authorization", "ApiKey my-api-key")
	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("Expected nil, got error: %v", err)
	}
	if apiKey != "my-api-key" {
		t.Errorf("Expected API key 'my-api-key', got '%s'", apiKey)
	}
}
