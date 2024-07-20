package controller

import (
	"net/http"
	"your-sales-report/internal/service"

	"github.com/labstack/echo/v4"
)

type ReportData struct {
	TotalGMV string
}

type ReportController struct {
	reportService *service.ReportService
}

func NewReportController(s *service.ReportService) *ReportController {
	return &ReportController{reportService: s}
}

func (h *ReportController) GetReports(c echo.Context) error {
	value, err := h.reportService.GetTotalGMV()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get total gmv"})

	}

	data := ReportData{
		TotalGMV: value.Value,
	}

	return c.Render(200, "index", data)
}
