package tasks

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/neptune2k21/chemin-d-or/internal/domain"
)

type postgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Create(ctx context.Context, req CreateTaskRequest) (*domain.Task, error) {
	const q = `
		INSERT INTO tasks (title, description, due_date)
		VALUES ($1, $2, $3)
		RETURNING id, title, description, due_date, created_at, updated_at`

	task := &domain.Task{}
	err := r.db.QueryRow(ctx, q, req.Title, req.Description, req.DueDate).
		Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("create task: %w", err)
	}
	return task, nil
}

func (r *postgresRepository) GetByID(ctx context.Context, id string) (*domain.Task, error) {
	const q = `
		SELECT id, title, description, due_date, created_at, updated_at
		FROM tasks WHERE id = $1`

	task := &domain.Task{}
	err := r.db.QueryRow(ctx, q, id).
		Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("get task by id: %w", err)
	}
	return task, nil
}

func (r *postgresRepository) List(ctx context.Context) ([]*domain.Task, error) {
	const q = `
		SELECT id, title, description, due_date, created_at, updated_at
		FROM tasks ORDER BY created_at DESC`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("list tasks: %w", err)
	}
	defer rows.Close()

	var tasks []*domain.Task
	for rows.Next() {
		task := &domain.Task{}
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan task: %w", err)
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *postgresRepository) Update(ctx context.Context, id string, req UpdateTaskRequest) (*domain.Task, error) {
	const q = `
		UPDATE tasks
		SET
			title       = COALESCE($1, title),
			description = COALESCE($2, description),
			due_date    = COALESCE($3, due_date),
			updated_at  = NOW()
		WHERE id = $4
		RETURNING id, title, description, due_date, created_at, updated_at`

	task := &domain.Task{}
	err := r.db.QueryRow(ctx, q, req.Title, req.Description, req.DueDate, id).
		Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("update task: %w", err)
	}
	return task, nil
}

func (r *postgresRepository) Delete(ctx context.Context, id string) error {
	const q = `DELETE FROM tasks WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	if err != nil {
		return fmt.Errorf("delete task: %w", err)
	}
	return nil
}
