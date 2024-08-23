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

type ReportsData struct {
	ID                    int    `json:"id"`
	IncidentDate          string `json:"incidentDate"`
	TypeOfReport          string `json:"typeOfReport"`
	ClientName            string `json:"clientName"`
	ClientSurname         string `json:"clientSurname"`
	ClientAddress         string `json:"clientAddress"`
	RespondingOfficerName string `json:"respondingOfficerName"`
	ResponderCallSign     string `json:"responderCallSign"`
	ResponderArrivalTime  string `json:"responderArrivalTime"`
	OperatorName          string `json:"operatorName"`
	OperatorPosition      string `json:"operatorPosition"`
	Report                string `json:"report"`
	UserId                string `json:"userId"`
}

type CreateReportHandler struct{}

// Handler for creating a report
func (h CreateReportHandler) HandleCreateReport(c echo.Context) error {
	supabaseClient, err := functions.CreateSupabaseClient()

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error from line 37 handleCreateReport handler")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()}) // return 500
	}

	ctx := context.Background()

	user, err := supabaseClient.Auth.User(ctx, config.GetUserToken())

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "An error occurred coould not fetch reports data"})
	}

	if user == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}

	userId := user.ID

	// Get form values
	incidentDate := c.FormValue("incidentDate")
	typeOfReport := c.FormValue("typeOfReport")
	clientName := c.FormValue("clientName")
	clientSurname := c.FormValue("clientSurname")
	clientAddress := c.FormValue("clientAddress")
	responderName := c.FormValue("responderName")
	responderTime := c.FormValue("responderTime")
	responderCallSign := c.FormValue("responderCallSign")
	operatorName := c.FormValue("operatorName")
	operatorPosition := c.FormValue("operatorPosition")
	report := c.FormValue("report")

	// Prepare query
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
		OperatorPosition:      operatorPosition,
		Report:                report,
		UserId:                userId,
	}

	var results []ReportsData
	err = supabaseClient.DB.From("reports").Insert(query).Execute(&results)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error from line 71 handleCreateReport handler")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()}) // return 400
	}

	return c.JSON(http.StatusOK, "Report created")

}

// Function to display create report form
func (h CreateReportHandler) HandleShowCreateReportForm(c echo.Context) error {
	supabaseClient, err := functions.CreateSupabaseClient()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "mhe"})
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

	return render(c, dashboard.CreateReportForm())
}
