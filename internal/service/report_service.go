package service

import "your-sales-report/internal/domain"

type ReportRepository interface {
	GetReports() ([]domain.Report, error)
}

type ReportService struct {
	reportRepository ReportRepository
}

func NewReportService(reportRepository ReportRepository) *ReportService {
	return &ReportService{
		reportRepository: reportRepository,
	}
}

func (s *ReportService) GetReports() ([]domain.Report, error) {
	return s.reportRepository.GetReports()
}
