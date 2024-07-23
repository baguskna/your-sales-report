package main

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"log"
	"your-sales-report/db"
	"your-sales-report/internal/controller"
	"your-sales-report/internal/repository"
	"your-sales-report/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
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

	// Service
	reportService := service.NewReportService(reportRepo)

	// Controller
	reportController := controller.NewReportController(reportService)

	e.GET("/", reportController.GetReports)
	e.GET("/ask_ai", controller.AskAIController)
	llm, err := openai.New()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	completion, err := llm.Call(ctx, "what companies does elon has",
		llms.WithTemperature(0.8),
		llms.WithStopWords([]string{""}),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(completion, "open ai")

	e.Logger.Fatal(e.Start(":3000"))
}
