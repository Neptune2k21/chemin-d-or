package tasks

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/neptune2k21/chemin-d-or/internal/domain"
)

var ErrTaskNotFound = errors.New("tâche non trouvée")

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, req CreateTaskRequest) (*domain.Task, error) {
	if req.Title == "" {
		return nil, errors.New("le titre est requis")
	}
	return s.repo.Create(ctx, req)
}

func (s *Service) GetByID(ctx context.Context, id string) (*domain.Task, error) {
	task, err := s.repo.GetByID(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrTaskNotFound
	}
	return task, err
}

func (s *Service) List(ctx context.Context) ([]*domain.Task, error) {
	return s.repo.List(ctx)
}

func (s *Service) Update(ctx context.Context, id string, req UpdateTaskRequest) (*domain.Task, error) {
	task, err := s.repo.Update(ctx, id, req)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrTaskNotFound
	}
	return task, err
}

func (s *Service) Delete(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrTaskNotFound
	}
	return err
}
