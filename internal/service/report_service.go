package service

import "your-sales-report/internal/domain"

type ReportRepository interface {
	GetReports() ([]domain.Report, error)
	GetTotalGMV() (*domain.GMV, error)
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

func (s *ReportService) GetTotalGMV() (*domain.GMV, error) {
	return s.reportRepository.GetTotalGMV()
}
