package handler

import (
	"fmt"
	"net/http"

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

	userId, err := config.GetUserIdFromCookie(c.Request(), c.Echo().AcquireContext(), supabaseClient)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error from line 47 handleCreateReport handler")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()}) // return 500
	}

	// Get form values
	incidentDate := c.FormValue("incidentDate")
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
	return render(c, dashboard.CreateReportForm())
}
