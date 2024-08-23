package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AskAIController(c echo.Context) error {
	return c.Render(200, "ask_ai", "")
}

func AskAIPostController(c echo.Context) error {
	query := c.FormValue("query")
	if query == "" {
		return c.String(http.StatusBadRequest, "query can't be null")
	}
	response := "okok " + query
	return c.String(http.StatusOK, response)
}
