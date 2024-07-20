package main

import (
	"html/template"
	"io"
	"your-sales-report/db"
	"your-sales-report/internal/controller"
	"your-sales-report/internal/repository"
	"your-sales-report/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}
func main() {
	db.InitDB()
	dbCon := db.GetDB()
	defer dbCon.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// serve static file in go
	e.Static("/images", "images")
	e.Static("/css", "css")

	e.Renderer = newTemplate()

	// Repository
	reportRepo := repository.NewReportRepository(dbCon)
	// gmvRepo := repository.NewGMVRepository(dbCon)

	// Service
	reportService := service.NewReportService(reportRepo)
	// gmvService := service.NewGMVService(gmvRepo)

	// Controller
	reportController := controller.NewReportController(reportService)
	// gmvController := controller.NewGMVController(gmvService)

	e.GET("/", reportController.GetReports)

	e.Logger.Fatal(e.Start(":3000"))
}
