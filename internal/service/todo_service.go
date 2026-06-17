package service

import (
	"time"
	"todo_cli/internal/domain"
	"todo_cli/internal/repository"

	"github.com/google/uuid"
)

type TodoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) *TodoService {
	return &TodoService{
		repo: repo,
	}
}

func (s *TodoService) Create(title string) error {
	now := time.Now()
	return s.repo.Create(
		domain.Todo{
			ID:        uuid.NewString(),
			Title:     title,
			Done:      false,
			CreatedAt: now,
			UpdatedAt: now,
		},
	)
}

func (s *TodoService) GetAll() ([]domain.Todo, error) {
	return s.repo.GetAll()
}

func (s *TodoService) Update(todo domain.Todo) error {
	todo.UpdatedAt = time.Now()

	return s.repo.Update(todo)
}

func (s *TodoService) Delete(id string) error {
	return s.repo.Delete(id)
}
