package handler

import (
	"github.com/kerneels94/reports/view/auth"
	"github.com/labstack/echo/v4"
)

type SignUpHandler struct{}

func (h SignUpHandler) HandleSignUp(c echo.Context) error {
	// userData := model.UserModel{
	// 	Email: "test@gmail.com",
	// }
	return render(c, auth.SignUpPage())
}
