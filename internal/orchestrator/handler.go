package orchestrator

import (
	"distributed-calculator/internal/models"
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Calculate(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Expression string `json:"expression"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusUnprocessableEntity)
		return
	}

	id := h.service.AddExpression(req.Expression)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func (h *Handler) GetExpressions(w http.ResponseWriter, r *http.Request) {
	expressions := h.service.GetExpressions()
	json.NewEncoder(w).Encode(map[string]interface{}{"expressions": expressions})
}

func (h *Handler) GetExpressionByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	expr, exists := h.service.GetExpressionByID(id)
	if !exists {
		http.Error(w, "Expression not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"expression": expr})
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	task, exists := h.service.GetTask()
	if !exists {
		http.Error(w, "No tasks available", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"task": task})
}

func (h *Handler) SubmitTaskResult(w http.ResponseWriter, r *http.Request) {
	var result models.TaskResult
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		http.Error(w, "Invalid request", http.StatusUnprocessableEntity)
		return
	}

	if !h.service.SubmitTaskResult(&result) {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
