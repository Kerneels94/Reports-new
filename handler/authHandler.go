package handler

import (
	"github.com/kerneels94/reports/view/auth"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct{}

func (h AuthHandler) HandleUserLogin(c echo.Context) error {
	// userData := model.UserModel{
	// 	Email: "test@gmail.com",
	// }
	return render(c, auth.LoginPage())
}
