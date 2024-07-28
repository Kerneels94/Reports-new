package handler

import (
	"github.com/kerneels94/reports/model"
	"github.com/kerneels94/reports/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	userData := model.UserModel{
		Email: "test@gmail.com",
	}
	return render(c, user.Show(userData))
}
