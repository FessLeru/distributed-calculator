package orchestrator

import (
	"distributed-calculator/internal/models"
	"sync"
	"time"
)

type Service struct {
	expressions map[string]*models.Expression
	tasks       []*models.Task
	mu          sync.Mutex
}

func NewService() *Service {
	return &Service{
		expressions: make(map[string]*models.Expression),
		tasks:       []*models.Task{},
	}
}

func (s *Service) AddExpression(expr string) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := generateID()
	s.expressions[id] = &models.Expression{
		ID:     id,
		Status: "pending",
		Result: 0,
	}

	tasks := parseExpression(expr)
	for _, task := range tasks {
		task.ID = generateID()
		s.tasks = append(s.tasks, task)
	}

	return id
}

func (s *Service) GetExpressions() map[string]*models.Expression {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.expressions
}

func (s *Service) GetExpressionByID(id string) (*models.Expression, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	expr, exists := s.expressions[id]
	return expr, exists
}

func (s *Service) GetTask() (*models.Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.tasks) == 0 {
		return nil, false
	}

	task := s.tasks[0]
	s.tasks = s.tasks[1:]
	return task, true
}

func (s *Service) SubmitTaskResult(result *models.TaskResult) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, expr := range s.expressions {
		if expr.Status == "pending" {
			expr.Result = result.Result
			expr.Status = "completed"
			return true
		}
	}
	return false
}

func generateID() string {
	return time.Now().Format("20060102150405")
}

func parseExpression(expr string) []*models.Task {
	return []*models.Task{
		{Arg1: 2, Arg2: 2, Operation: "+", OperationTime: 1000},
		{Arg1: 4, Arg2: 2, Operation: "*", OperationTime: 2000},
	}
}
