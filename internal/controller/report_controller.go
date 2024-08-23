package controller

import (
	"bytes"
	"html/template"
	"net/http"
	"your-sales-report/internal/domain"
	"your-sales-report/internal/service"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/labstack/echo/v4"
)

type ReportData struct {
	TotalGMV     string
	GMVChartHTML template.HTML
}

type ReportController struct {
	reportService *service.ReportService
}

func NewReportController(s *service.ReportService) *ReportController {
	return &ReportController{reportService: s}
}

func (h *ReportController) GetReports(c echo.Context) error {
	valueGMV, err := h.reportService.GetTotalGMV()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get total gmv"})

	}

	orders, err := h.reportService.GetTotalOrderAndPercentageByMarketplace()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get total order and percentage"})
	}

	data := ReportData{
		TotalGMV:     valueGMV.Value,
		GMVChartHTML: generateTotalOrderChart(orders),
	}

	return c.Render(200, "index", data)
}

func generateTotalOrderChart(orders []domain.TotalOrderAndPercentage) template.HTML {
	bar := charts.NewBar()

	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Total Order"}),
	)

	xAxis := make([]string, len(orders))
	yAxis := make([]opts.BarData, len(orders))

	for i, order := range orders {
		xAxis[i] = order.Marketplace
		yAxis[i] = opts.BarData{Value: order.TotalOrder}
	}

	bar.SetXAxis(xAxis).
		AddSeries("Total Order", yAxis)

	var buf bytes.Buffer
	if err := bar.Render(&buf); err != nil {
		panic(err)
	}

	return template.HTML(buf.String())
}
