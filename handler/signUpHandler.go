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
	return Render(c, auth.SignUpPage())
}

func (h SignUpHandler) HandleUserSignUp(c echo.Context) error {
	email := c.FormValue("email")
	firstName := c.FormValue("firstName")
	lastName := c.FormValue("lastName")
	password := c.FormValue("password")

	if email == "" || password == "" || len(email) <= 0 || len(password) <= 0 || firstName == "" || len(firstName) <= 0 || lastName == "" || len(lastName) <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Please check your credentials"})
	}

	supabaseClient, err := functions.CreateSupabaseClient()

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	ctx := context.Background()

	_, err = supabaseClient.Auth.SignUp(ctx, supa.UserCredentials{
		Email:    email,
		Password: password,
		Data: map[string]interface{}{
			"firstName": firstName,
			"lastName":  lastName,
		},
	})

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "An error while signing up but you won't know what it is! because it's golang"})
	}

	functions.HtmxRedirect(c, "/login")

	return c.JSON(http.StatusOK, "Signed up")
}
