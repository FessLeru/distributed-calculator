package main

import (
	"bytes"
	"distributed-calculator/internal/models"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var orchestratorURL = "http://localhost:8080"

func main() {
	computingPower := 4 // Количество горутин
	for i := 0; i < computingPower; i++ {
		go worker()
	}

	select {}
}

func worker() {
	for {
		task, err := fetchTask()
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}

		result := executeTask(task)
		if err := submitTaskResult(task.ID, result); err != nil {
			log.Println("Failed to submit result:", err)
		}
	}
}

func fetchTask() (*models.Task, error) {
	resp, err := http.Get(orchestratorURL + "/internal/task")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Task *models.Task `json:"task"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Task, nil
}

func executeTask(task *models.Task) float64 {
	time.Sleep(time.Duration(task.OperationTime) * time.Millisecond)

	switch task.Operation {
	case "+":
		return task.Arg1 + task.Arg2
	case "-":
		return task.Arg1 - task.Arg2
	case "*":
		return task.Arg1 * task.Arg2
	case "/":
		return task.Arg1 / task.Arg2
	default:
		return 0
	}
}

func submitTaskResult(taskID string, result float64) error {
	data := models.TaskResult{
		ID:     taskID,
		Result: result,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = http.Post(orchestratorURL+"/internal/task", "application/json", bytes.NewBuffer(jsonData))
	return err
}
