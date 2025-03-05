package orchestrator

import (
	"bytes"
	"distributed-calculator/internal/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Calculate(t *testing.T) {
	service := NewService()
	handler := NewHandler(service)

	body := bytes.NewBufferString(`{"expression": "2+2*2"}`)
	req := httptest.NewRequest("POST", "/api/v1/calculate", body)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handler.Calculate(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, rec.Code)
	}

	var response struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}
	if response.ID == "" {
		t.Error("Expected non-empty ID, got empty string")
	}
}

func TestHandler_GetExpressions(t *testing.T) {
	service := NewService()
	handler := NewHandler(service)

	service.AddExpression("2+2*2")

	req := httptest.NewRequest("GET", "/api/v1/expressions", nil)
	rec := httptest.NewRecorder()

	handler.GetExpressions(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	var response struct {
		Expressions []*models.Expression `json:"expressions"`
	}
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}
	if len(response.Expressions) == 0 {
		t.Error("Expected at least one expression, got none")
	}
}
