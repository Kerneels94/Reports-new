package handler

import (
	"context"
	"fmt"
	"net/http"

	supa "github.com/nedpals/supabase-go"

	"github.com/kerneels94/reports/config"
	"github.com/kerneels94/reports/functions"
	"github.com/kerneels94/reports/view/dashboard"
	"github.com/labstack/echo/v4"
)

type DashboardHandler struct{}

func (h DashboardHandler) HandleDashboard(c echo.Context) error {
	return Render(c, dashboard.DashboardPage())
}

func (h DashboardHandler) HandleLogout(c echo.Context) error {
	config.CookieLogout(c.Response())

	functions.HtmxRedirect(c, "/login")

	return c.JSON(http.StatusOK, map[string]string{"message": "Logged out successfully"})
}

// Dashboard - Users
func (h DashboardHandler) HandleUsers(c echo.Context) error {
	return Render(c, dashboard.DashboardUsersPage())
}

// Dashboard - Users - Add User
func (h DashboardHandler) HandleAddUser(c echo.Context) error {
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

	// Add user
	_, err = supabaseClient.Auth.SignUp(ctx, supa.UserCredentials{
		Email:    email,
		Password: password,
	})

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "An while adding user"})
	}
	// todo find another way todo this
	row := User{
		ID:        "todo",
		FirstName: c.FormValue("name"),
		LastName:  c.FormValue("surname"),
		Role:      "admin",
	}

	// Add data to the user table
	var results []User
	err = supabaseClient.DB.From("users").Insert(row).Execute(&results)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "An error occurred while creating user"})
	}

	return c.JSON(http.StatusOK, "Account added")
}

// Dashboard - Users - Get All Users
func (h DashboardHandler) HandleGetAllUser(c echo.Context) ([]dashboard.User, error) {
	supabaseClient, err := functions.CreateSupabaseClient()

	if err != nil {
		fmt.Println(err)
		fmt.Print("Error HandleGetAllUser Line 147")
		return nil, c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	var results []map[string]interface{}

	// Get all users
	err = supabaseClient.DB.From("users").Select("*").Execute(&results)
	if err != nil {
		fmt.Println(err)
		fmt.Print("Error HandleGetAllUser Line 155")
		return nil, c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Format the results
	var users []dashboard.User
	for _, result := range results {
		user := dashboard.User{
			Email: result["email"].(string),
		}
		users = append(users, user)
	}

	return users, nil
}

// Dashboard - Users - Table Page
func (h DashboardHandler) HandleDashboardUsersTablePage(c echo.Context) error {
	users, err := h.HandleGetAllUser(c)

	if err != nil {
		fmt.Println(err)
		fmt.Print("Error HandleGetAllUser Line 169")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return Render(c, dashboard.DashboardUserTablePage(users))
}
