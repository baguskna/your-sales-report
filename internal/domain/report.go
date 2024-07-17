package domain

import "github.com/google/uuid"

type Report struct {
	ID   uuid.UUID
	Name string
}

type Stat struct {
	ID       uuid.UUID
	Total    int
	Date     string
	ReportID uuid.UUID
}
