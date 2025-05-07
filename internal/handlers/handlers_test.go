package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	HealthCheckHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]string
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response["status"] != "healthy" {
		t.Errorf("Expected status 'healthy', got %s", response["status"])
	}
}

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/hello", nil)
	w := httptest.NewRecorder()

	HelloHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]string
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response["message"] != "Hello, World!" {
		t.Errorf("Expected message 'Hello, World!', got %s", response["message"])
	}
}

func TestEchoHandler(t *testing.T) {
	testBody := "Hello, Server!"
	req := httptest.NewRequest(http.MethodPost, "/api/v1/echo", bytes.NewBufferString(testBody))
	w := httptest.NewRecorder()

	EchoHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response["message"] != "Echo" {
		t.Errorf("Expected message 'Echo', got %s", response["message"])
	}
	if response["data"] != testBody {
		t.Errorf("Expected data '%s', got %s", testBody, response["data"])
	}
}

func TestMethodNotAllowed(t *testing.T) {
	tests := []struct {
		name       string
		handler    http.HandlerFunc
		method     string
		path       string
		body       string
		wantStatus int
	}{
		{
			name:       "HealthCheck POST",
			handler:    HealthCheckHandler,
			method:     http.MethodPost,
			path:       "/health",
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			name:       "Hello POST",
			handler:    HelloHandler,
			method:     http.MethodPost,
			path:       "/api/v1/hello",
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			name:       "Echo GET",
			handler:    EchoHandler,
			method:     http.MethodGet,
			path:       "/api/v1/echo",
			wantStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, bytes.NewBufferString(tt.body))
			w := httptest.NewRecorder()

			tt.handler(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, w.Code)
			}
		})
	}
} 