package domain

import (
	"github.com/google/uuid"
)

type Report struct {
	ID    uuid.UUID
	Title string
	Stats []Stat
}

type Stat struct {
	ID       uuid.UUID
	Total    int
	Date     string
	ReportID uuid.UUID // foreign key of report id
}
