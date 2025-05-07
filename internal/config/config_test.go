package config

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	// Test with default values
	cfg := New()
	if cfg.Port != 8080 {
		t.Errorf("Expected default port 8080, got %d", cfg.Port)
	}
	if cfg.ShutdownTimeout != 10 {
		t.Errorf("Expected default shutdown timeout 10, got %d", cfg.ShutdownTimeout)
	}
	if cfg.Environment != "development" {
		t.Errorf("Expected default environment 'development', got %s", cfg.Environment)
	}

	// Test with custom values
	os.Setenv("PORT", "9090")
	os.Setenv("SHUTDOWN_TIMEOUT", "20")
	os.Setenv("ENV", "production")
	defer func() {
		os.Unsetenv("PORT")
		os.Unsetenv("SHUTDOWN_TIMEOUT")
		os.Unsetenv("ENV")
	}()

	cfg = New()
	if cfg.Port != 9090 {
		t.Errorf("Expected port 9090, got %d", cfg.Port)
	}
	if cfg.ShutdownTimeout != 20 {
		t.Errorf("Expected shutdown timeout 20, got %d", cfg.ShutdownTimeout)
	}
	if cfg.Environment != "production" {
		t.Errorf("Expected environment 'production', got %s", cfg.Environment)
	}
}

func TestGetEnv(t *testing.T) {
	// Test with existing environment variable
	os.Setenv("TEST_KEY", "test_value")
	defer os.Unsetenv("TEST_KEY")

	value := getEnv("TEST_KEY", "default")
	if value != "test_value" {
		t.Errorf("Expected 'test_value', got %s", value)
	}

	// Test with non-existing environment variable
	value = getEnv("NON_EXISTING_KEY", "default")
	if value != "default" {
		t.Errorf("Expected 'default', got %s", value)
	}
} 