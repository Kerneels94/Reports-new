package handler

import (
	"context"
	"fmt"
	"net/http"

	supa "github.com/nedpals/supabase-go"

	"github.com/kerneels94/reports/functions"
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

func (h LoginHandler) HandleUserLoginLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Email and password are required"})
	}

	supabaseClient, err := functions.CreateSupabaseClient()

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	ctx := context.Background()

	user, err := supabaseClient.Auth.SignIn(ctx, supa.UserCredentials{
		Email:    email,
		Password: password,
	})

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}
