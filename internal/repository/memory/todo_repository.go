package memory

import (
	"todo_cli/internal/domain"
)

type TodoRepository struct {
	items []domain.Todo
}

func New() *TodoRepository {
	return &TodoRepository{
		items: []domain.Todo{},
	}
}

func (r *TodoRepository) GetAll() ([]domain.Todo, error) {
	return r.items, nil
}

func (r *TodoRepository) Create(todo domain.Todo) error {
	r.items = append(r.items, todo)
	return nil
}

func (r *TodoRepository) Update(todo domain.Todo) error {
	return nil
}

func (r *TodoRepository) Delete(id string) error {
	return nil
}