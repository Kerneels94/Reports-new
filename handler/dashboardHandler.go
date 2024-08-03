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

type ReportsData struct {
	ID                    int
	IncidentDate          string
	TypeOfReport          string
	ClientName            string
	ClientSurname         string
	ClientAddress         string
	RespondingOfficerName string
	ResponderCallSign     string
	ResponderArrivalTime  string
	OperatorName          string
	OperatorRank          string
	Report                string
}

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

/*
	Handler for creating a report
*/

func (h DashboardHandler) HandleCreateReport(c echo.Context) error {
	supabaseClient, err := functions.CreateSupabaseClient()

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	incidentDate := c.FormValue("incidentDate")
	typeOfReport := c.FormValue("typeOfReport")
	clientName := c.FormValue("clientName")
	clientSurname := c.FormValue("clientSurname")
	clientAddress := c.FormValue("clientAddress")
	responderName := c.FormValue("responderName")
	responderTime := c.FormValue("responderTime")
	responderCallSign := c.FormValue("responderCallSign")
	operatorName := c.FormValue("operatorName")
	operatorRank := c.FormValue("operatorRank")
	report := c.FormValue("report")

	query := ReportsData{
		IncidentDate:          incidentDate,
		TypeOfReport:          typeOfReport,
		ClientName:            clientName,
		ClientAddress:         clientAddress,
		ClientSurname:         clientSurname,
		RespondingOfficerName: responderName,
		ResponderCallSign:     responderCallSign,
		ResponderArrivalTime:  responderTime,
		OperatorName:          operatorName,
		OperatorRank:          operatorRank,
		Report:                report,
	}

	var results []ReportsData
	err = supabaseClient.DB.From("reports").Insert(query).Execute(&results)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, "Report created")

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
