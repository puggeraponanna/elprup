package models

import (
	"time"
)

type Module[T any] struct {
	Slug        string
	Name        string
	Description string
	Schema      T
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	CreatedBy   string
	UpdatedBy   string
}
