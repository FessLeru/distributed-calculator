package main

import (
	"distributed-calculator/internal/orchestrator"
	"log"
	"net/http"
)

func main() {
	service := orchestrator.NewService()
	handler := orchestrator.NewHandler(service)

	http.HandleFunc("/api/v1/calculate", handler.Calculate)
	http.HandleFunc("/api/v1/expressions", handler.GetExpressions)
	http.HandleFunc("/api/v1/expressions/", handler.GetExpressionByID)
	http.HandleFunc("/internal/task", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.GetTask(w, r)
		} else if r.Method == http.MethodPost {
			handler.SubmitTaskResult(w, r)
		}
	})

	log.Println("Orchestrator is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
