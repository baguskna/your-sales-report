package controller

import "github.com/labstack/echo/v4"

func AskAIController(c echo.Context) error {
	return c.Render(200, "ask_ai", "")
}
