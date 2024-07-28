package handler

import (
	"github.com/kerneels94/reports/view/auth"
	"github.com/labstack/echo/v4"
)

type LoginHandler struct{}

func (h LoginHandler) HandleUserLogin(c echo.Context) error {
	// userData := model.UserModel{
	// 	Email: "test@gmail.com",
	// }
	return render(c, auth.LoginPage())
}
