package repository

import (
	"database/sql"
	"your-sales-report/internal/domain"
)

type ReportRepository struct {
	DB *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{DB: db}
}

func (r *ReportRepository) GetReports() ([]domain.Report, error) {
	query := "SELECT id, total, date FROM reports"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []domain.Report

	for rows.Next() {
		var report domain.Report
		if err := rows.Scan(&report.ID, &report.Total, &report.Date); err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}
