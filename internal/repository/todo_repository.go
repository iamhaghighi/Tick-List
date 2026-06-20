package repository

import (
	"context"
	"todo_cli/internal/domain"
)

type TodoRepository interface {
	GetAll(ctx context.Context) ([]domain.Todo, error)
	Create(ctx context.Context, todo domain.Todo) (domain.Todo, error)
	Update(ctx context.Context, todo domain.Todo) error
	Delete(ctx context.Context, id int64) error
	Toggle(ctx context.Context, id int64, done bool) error
}
