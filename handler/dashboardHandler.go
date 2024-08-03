package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kerneels94/reports/config"
	"github.com/kerneels94/reports/functions"
	"github.com/kerneels94/reports/view/dashboard"
	"github.com/labstack/echo/v4"
)

type DashboardHandler struct{}

func (h DashboardHandler) HandleDashboard(c echo.Context) error {
	supabaseClient, err := functions.CreateSupabaseClient()

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	ctx := context.Background()

	user, err := supabaseClient.Auth.User(ctx, config.GetUserToken())

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
	supabaseClient, err := functions.CreateSupabaseClient()

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	ctx := context.Background()

	err = supabaseClient.Auth.SignOut(ctx, config.GetUserToken())

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "An error occurred while logging out"})
	}

	config.SetUserToken("")

	functions.HtmxRedirect(c, "/login")

	return c.JSON(http.StatusOK, map[string]string{"message": "Logged out successfully"})
}
