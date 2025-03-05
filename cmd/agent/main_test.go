package main

import (
	"distributed-calculator/internal/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MockOrchestrator() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/internal/task":
			if r.Method == http.MethodGet {
				task := models.Task{
					ID:            "1",
					Arg1:          2,
					Arg2:          2,
					Operation:     "+",
					OperationTime: 1000,
				}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(map[string]interface{}{"task": task})
			} else if r.Method == http.MethodPost {
				var result models.TaskResult
				if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
					w.WriteHeader(http.StatusUnprocessableEntity)
					return
				}
				w.WriteHeader(http.StatusOK)
			}
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
}

func TestFetchTask(t *testing.T) {
	server := MockOrchestrator()
	defer server.Close()

	oldURL := orchestratorURL
	orchestratorURL = server.URL
	defer func() { orchestratorURL = oldURL }()

	task, err := fetchTask()
	if err != nil {
		t.Fatalf("Failed to fetch task: %v", err)
	}

	if task.ID != "1" {
		t.Errorf("Expected task ID '1', got '%s'", task.ID)
	}
	if task.Arg1 != 2 || task.Arg2 != 2 || task.Operation != "+" {
		t.Errorf("Expected task 2 + 2, got %f %s %f", task.Arg1, task.Operation, task.Arg2)
	}
}

func TestExecuteTask(t *testing.T) {
	tests := []struct {
		name     string
		task     models.Task
		expected float64
	}{
		{"addition", models.Task{Arg1: 2, Arg2: 2, Operation: "+"}, 4},
		{"subtraction", models.Task{Arg1: 5, Arg2: 3, Operation: "-"}, 2},
		{"multiplication", models.Task{Arg1: 3, Arg2: 3, Operation: "*"}, 9},
		{"division", models.Task{Arg1: 6, Arg2: 2, Operation: "/"}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := executeTask(&tt.task)
			if result != tt.expected {
				t.Errorf("Expected %f, got %f", tt.expected, result)
			}
		})
	}
}

func TestSubmitTaskResult(t *testing.T) {
	server := MockOrchestrator()
	defer server.Close()

	oldURL := orchestratorURL
	orchestratorURL = server.URL
	defer func() { orchestratorURL = oldURL }()

	err := submitTaskResult("1", 4.0)
	if err != nil {
		t.Fatalf("Failed to submit task result: %v", err)
	}
}
