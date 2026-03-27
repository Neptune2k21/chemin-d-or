package tasks

import (
	"context"

	"github.com/neptune2k21/chemin-d-or/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, req CreateTaskRequest) (*domain.Task, error)
	GetByID(ctx context.Context, id string) (*domain.Task, error)
	List(ctx context.Context) ([]*domain.Task, error)
	Update(ctx context.Context, id string, req UpdateTaskRequest) (*domain.Task, error)
	Delete(ctx context.Context, id string) error
}
