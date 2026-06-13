package domain

import "time"

type Todo struct {
	ID string
	Title string
	Done bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
