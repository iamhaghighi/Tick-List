package domain

import "time"

type Todo struct {
	ID        int64
	Title     string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time

	// CompletedAt time.Time

	// Description string

	// Tags []string
}
