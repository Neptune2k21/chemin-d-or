package domain

import "time"

type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description,omitempty"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
