package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AskAIController(c echo.Context) error {
	return c.Render(200, "ask_ai", "")
}

func AskAIPostController(c echo.Context) error {
	query := c.FormValue("query")
	fmt.Println(query, "controller ai")
	response := "okok " + query
	return c.String(http.StatusOK, response)
}
