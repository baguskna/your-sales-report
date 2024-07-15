package main

import (
	"html/template"
	"io"

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

type Stats struct {
	Total int
	Date  string
}

type Data struct {
	Stats []Stats
	Name  string
}

func newStats(total int, date string) Stats {
	return Stats{
		Total: total,
		Date:  date,
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	// serve static file in go
	e.Static("/images", "images")
	e.Static("/css", "css")

	data := Data{
		Name: "Tanamera Coffee Drip Bag / Filter Bag: Breakfast Blend",
		Stats: []Stats{
			newStats(4126, "Sun 14 Jul 2024"),
		},
	}

	e.Renderer = newTemplate()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", data)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
