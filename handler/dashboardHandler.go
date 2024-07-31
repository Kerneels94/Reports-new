package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"

	supa "github.com/nedpals/supabase-go"

	"github.com/kerneels94/reports/view/dashboard"
	"github.com/labstack/echo/v4"
)

type DashboardHandler struct{}

func (h DashboardHandler) HandleDashboard(c echo.Context) error {
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

	// todo - set access_token from login -> user SetUserToken from the config file 
	user, err := supabaseClient.Auth.User(ctx, "")

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "An error occurred or you are not logged in."})
	}

	if user == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}

	fmt.Println(user)

	return render(c, dashboard.DashboardPage())
}

func (h DashboardHandler) HandleLogout(c echo.Context) error {
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

	// todo - set access_token from login -> user SetUserToken from the config file 
	err := supabaseClient.Auth.SignOut(ctx, "")

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "An error occurred while logging out"})
	}

	c.Response().Header().Set("HX-Redirect", "/login")
    return c.JSON(http.StatusOK, map[string]string{"message": "Logged out successfully"})
}
