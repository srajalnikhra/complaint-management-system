package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestSuccessResponse verifies that Success writes
// the expected JSON response and status code.
func TestSuccessResponse(t *testing.T) {
	rec := httptest.NewRecorder()

	Success(rec, http.StatusOK, "success", map[string]string{
		"name": "Srajal",
	})

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var response APIResponse

	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if !response.Success {
		t.Error("expected success to be true")
	}

	if response.Message != "success" {
		t.Errorf("expected message success, got %s", response.Message)
	}
}

// TestErrorResponse verifies that Error writes
// the expected JSON response and status code.
func TestErrorResponse(t *testing.T) {
	rec := httptest.NewRecorder()

	Error(rec, http.StatusBadRequest, "invalid request")

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}

	var response APIResponse

	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if response.Success {
		t.Error("expected success to be false")
	}

	if response.Message != "invalid request" {
		t.Errorf("expected invalid request, got %s", response.Message)
	}
}