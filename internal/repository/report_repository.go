package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"
	"your-sales-report/internal/domain"

	"github.com/google/uuid"
)

type ReportRepository struct {
	DB *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{DB: db}
}

func (r *ReportRepository) GetReports() ([]domain.Report, error) {
	query := `
	SELECT 
		r.id AS reports_id, r.title, 
		s.id AS stat_id, s.total, s.date
	FROM 
		reports r
	LEFT JOIN 
		stats s ON r.id = s.reports_id
	ORDER BY s.date ASC`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fmt.Println(query)

	reportMap := make(map[uuid.UUID]*domain.Report)
	for rows.Next() {
		var reportID, statID uuid.UUID
		var reportTitle string
		var statDate time.Time
		var statTotal int

		if err := rows.Scan(&reportID, &reportTitle, &statID, &statTotal, &statDate); err != nil {
			return nil, err
		}

		report, exists := reportMap[reportID]
		if !exists {
			report = &domain.Report{
				ID:    reportID,
				Title: reportTitle,
				Stats: []domain.Stat{},
			}
			reportMap[reportID] = report
		}

		if statID != uuid.Nil {
			statDateStr := statDate.Format("2006-01-02")
			stat := domain.Stat{
				ID:       statID,
				Total:    statTotal,
				Date:     statDateStr,
				ReportID: reportID,
			}
			report.Stats = append(report.Stats, stat)
		}
	}

	var reports []domain.Report
	for _, report := range reportMap {
		reports = append(reports, *report)
	}

	return reports, nil
}

func (r *ReportRepository) GetTotalGMV() (*domain.GMV, error) {
	query := `SELECT COALESCE(SUM(value), 0) as total_gmv FROM raw_data GROUP BY official_store;`

	var gmvValue float64

	err := r.DB.QueryRow(query).Scan(&gmvValue)
	if err != nil {
		log.Printf("Error fetching total GMV: %v", err)
		return &domain.GMV{Value: "0"}, err
	}

	fmt.Println(query)

	// Convert float64 to string with specific precision
	formattedValue := strconv.FormatFloat(gmvValue, 'f', -1, 64)

	return &domain.GMV{
		Value: formattedValue,
	}, err
}
