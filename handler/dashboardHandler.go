package handler

import (
	// "context"
	// "fmt"
	// "net/http"
	// "os"

	// supa "github.com/nedpals/supabase-go"

	"github.com/kerneels94/reports/view/dashboard"
	"github.com/labstack/echo/v4"
)

type DashboardHandler struct{}

func (h DashboardHandler) HandleDashboard(c echo.Context) error {
	return render(c, dashboard.DashboardPage())
}
