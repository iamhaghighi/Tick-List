package service

import (
	"context"
	"errors"
	"time"
	"todo_cli/internal/domain"
	"todo_cli/internal/repository"
)

type TodoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) Create(ctx context.Context, title string) (domain.Todo, error) {
	if title == "" {
		return domain.Todo{}, errors.New("title cannot be empty")
	}

	todo := domain.Todo{
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	created, err := s.repo.Create(ctx, todo)
	if err != nil {
		return domain.Todo{}, err
	}
	return created, nil
}

func (s *TodoService) GetAll(ctx context.Context) ([]domain.Todo, error) {
	return s.repo.GetAll(ctx)
}

func (s *TodoService) Update(ctx context.Context, todo domain.Todo) (domain.Todo, error) {
	if todo.ID == 0 {
		return domain.Todo{}, errors.New("invalid id")
	}
	todo.UpdatedAt = time.Now()
	if err := s.repo.Update(ctx, todo); err != nil {
		return domain.Todo{}, err
	}
	return todo, nil
}

func (s *TodoService) Delete(ctx context.Context, id int64) error {
	if id == 0 {
		return errors.New("invalid id")
	}
	return s.repo.Delete(ctx, id)
}

func (s *TodoService) UpdateDoneToggle() {
	
}