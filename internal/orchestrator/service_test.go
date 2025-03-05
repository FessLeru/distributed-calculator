package orchestrator

import (
	"distributed-calculator/internal/models"
	"testing"
)

func TestService_AddExpression(t *testing.T) {
	service := NewService()

	id := service.AddExpression("2+2*2")

	if id == "" {
		t.Error("Expected non-empty ID, got empty string")
	}

	expr, exists := service.GetExpressionByID(id)
	if !exists {
		t.Errorf("Expected expression with ID %s to exist, but it doesn't", id)
	}

	if expr.Status != "pending" {
		t.Errorf("Expected status 'pending', got '%s'", expr.Status)
	}
}

func TestService_GetTask(t *testing.T) {
	service := NewService()

	service.AddExpression("2+2*2")

	task, exists := service.GetTask()
	if !exists {
		t.Error("Expected task to exist, but it doesn't")
	}

	if task.Operation == "" {
		t.Error("Expected task to have an operation, got empty string")
	}
}

func TestService_SubmitTaskResult(t *testing.T) {
	service := NewService()

	id := service.AddExpression("2+2*2")

	task, _ := service.GetTask()

	result := &models.TaskResult{
		ID:     task.ID,
		Result: 4.0,
	}
	success := service.SubmitTaskResult(result)
	if !success {
		t.Error("Expected task result to be submitted successfully, but it wasn't")
	}

	expr, _ := service.GetExpressionByID(id)
	if expr.Result != 4.0 {
		t.Errorf("Expected expression result to be 4.0, got %f", expr.Result)
	}
}
