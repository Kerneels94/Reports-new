package handler

import (
	"github.com/kerneels94/reports/view/home"
	"github.com/labstack/echo/v4"
)

type MainPageHandler struct{}

func (h MainPageHandler) HandleShowMainPage(c echo.Context) error {
	return render(c, home.HomePage())
}
