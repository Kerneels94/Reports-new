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

type SignUpHandler struct{}

func (h SignUpHandler) HandleSignUp(c echo.Context) error {
	return render(c, auth.SignUpPage())
}

type User struct {
	ID         int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role       string `json:"role"`
}

func (h SignUpHandler) HandleUserSignUp(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || password == "" || len(email) <= 0 || len(password) <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Email and password are required"})
	}

	supabaseClient, err := functions.CreateSupabaseClient()

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	ctx := context.Background()

	user, err = supabaseClient.Auth.SignUp(ctx, supa.UserCredentials{
		Email:    email,
		Password: password,
	})

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "An error occurred"})
	}

	row := User{
		ID: user.ID,
		FirstName: c.FormValue("name"),
		LastName:  c.FormValue("surname"),
		Role:       "admin",
	}

	// Add data to the user table
	var results []User
	err = supabaseClient.DB.From("users").Insert(row).Execute(&results)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "An error occurred while creating user"})
	}

	functions.HtmxRedirect(c, "/login")

	return c.JSON(http.StatusOK, "Signed up")
}
