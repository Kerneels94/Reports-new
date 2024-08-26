package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/kerneels94/reports/config"
	"github.com/kerneels94/reports/functions"
	"github.com/kerneels94/reports/view/dashboard"
	"github.com/labstack/echo/v4"
)

type ReportsData struct {
	// ID                    int    `json:"id"`
	IncidentDate string `json:"incident_date"`
	// TypeOfReport          string `json:"type_of_report"`
	ClientName    string `json:"client_name"`
	ClientSurname string `json:"client_surname"`
	ClientAddress string `json:"client_address"`
	// RespondingOfficerName string `json:"responding_officer_name"`
	// ResponderCallSign     string `json:"responder_call_sign"`
	// ResponderArrivalTime  string `json:"responder_arrival_time"`
	OperatorName     string `json:"operator_name"`
	OperatorPosition string `json:"operator_position"`
	Report           string `json:"report"`
	UserId           string `json:"user_id"`
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

	// Get the date from the form
	// Get the date from the form
	incidentDate := c.FormValue("incidentDate") // This should be in a string format like "2024-08-26"

	// Parse the date string into a time.Time object
	parsedDate, err := time.Parse("2006-01-02", incidentDate)
	if err != nil {
		fmt.Println("Date could not be parsed. Check the format. It should be in the format YYYY-MM-DD.", err)
	}

	// Format the parsed date to a string that matches the database timestamp format
	formattedDate := parsedDate.Format("2006-01-02 15:04:05")

	// Pass `formattedDate` to the database
	fmt.Println("Formatted Date for DB:", formattedDate)

	// incidentDate := c.FormValue("incidentDate")
	// typeOfReport := c.FormValue("typeOfReport")
	clientName := c.FormValue("clientName")
	clientSurname := c.FormValue("clientSurname")
	clientAddress := c.FormValue("clientAddress")
	// responderName := c.FormValue("responderName")
	// responderTime := c.FormValue("responderTime")
	// responderCallSign := c.FormValue("responderCallSign")
	operatorName := c.FormValue("operatorName")
	operatorPosition := c.FormValue("operatorPosition")
	report := c.FormValue("report")

	// Prepare query
	query := ReportsData{
		IncidentDate: incidentDate,
		// TypeOfReport:          typeOfReport,
		ClientName:    clientName,
		ClientAddress: clientAddress,
		ClientSurname: clientSurname,
		// RespondingOfficerName: responderName,
		// ResponderCallSign:     responderCallSign,
		// ResponderArrivalTime:  responderTime,
		OperatorName:     operatorName,
		OperatorPosition: operatorPosition,
		Report:           report,
		UserId:           userId,
	}

	var results []ReportsData
	err = supabaseClient.DB.From("cake").Insert(query).Execute(&results)

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
