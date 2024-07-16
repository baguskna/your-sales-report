package domain

import "github.com/google/uuid"

type Report struct {
	ID    uuid.UUID
	Total int
	Date  string
}
