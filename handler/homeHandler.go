package handler

import (
	"github.com/kerneels94/reports/view/home"
	"github.com/labstack/echo/v4"
)

type MainPageHandler struct{}

func (h MainPageHandler) HandleShowMainPage(c echo.Context) error {
	// userData := model.UserModel{
	// 	Email: "test@gmail.com",
	// }
	return render(c, home.HomePage())
}
