package repository

import "todo_cli/internal/domain"

type TodoRepository interface {
	GetAll() ([]domain.Todo, error)
	Create(todo domain.Todo) error
	Update(todo domain.Todo) error
	Delete(id string) error
}
