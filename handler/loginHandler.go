package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"

	supa "github.com/nedpals/supabase-go"

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

	// API_URL
	API_URL := os.Getenv("API_URL")
	if API_URL == "" {
		fmt.Println("API_URL is not set")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "API_URL is not set"})
	}

	// API_KEY
	API_KEY := os.Getenv("API_KEY")
	if API_KEY == "" {
		fmt.Println("API_KEY is not set")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "API_KEY is not set"})
	}

	supabaseClient := supa.CreateClient(API_URL, API_KEY)

	ctx := context.Background()

	user, err := supabaseClient.Auth.SignIn(ctx, supa.UserCredentials{
		Email:    email,
		Password: password,
	})

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "An error occurred"})
	}

	return c.JSON(http.StatusOK, user)
}
