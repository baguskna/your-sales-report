package controller

import (
	"net/http"
	"your-sales-report/internal/service"

	"github.com/labstack/echo/v4"
)

type ReportController struct {
	reportService *service.ReportService
}

func NewReportController(s *service.ReportService) *ReportController {
	return &ReportController{reportService: s}
}

func (h *ReportController) GetReports(c echo.Context) error {
	reports, err := h.reportService.GetReports()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get reports"})
	}

	return c.Render(200, "index", reports)
}
