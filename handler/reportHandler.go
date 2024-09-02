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
	ClientCode            string `json:"client_code"`
	IncidentDate          string `json:"incident_date"`
	TypeOfReport          string `json:"report_type"`
	ClientName            string `json:"client_name"`
	ClientSurname         string `json:"client_surname"`
	ClientAddress         string `json:"client_address"`
	RespondingOfficerName string `json:"armed_response_officer_name"`
	ResponderCallSign     string `json:"armed_response_call_sign"`
	ResponderArrivalTime  string `json:"armed_response_arrival_time"`
	OperatorName          string `json:"operator_name"`
	OperatorPosition      string `json:"operator_position"`
	Report                string `json:"report"`
	UserId                string `json:"user_id"`
}

type ReportHandler struct{}

// Handler for creating a report
func (h ReportHandler) HandleCreateReport(c echo.Context) error {
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
	clientCode := c.FormValue("clientCode")
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
		ClientCode:            clientCode,
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
	err = supabaseClient.DB.From("cake").Insert(query).Execute(&results)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error from line 71 handleCreateReport handler")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()}) // return 400
	}

	return c.JSON(http.StatusOK, "Report created")

}

// Function to display create report form
func (h ReportHandler) HandleShowCreateReportForm(c echo.Context) error {
	return render(c, dashboard.CreateReportForm())
}

// Dashboard - Users - Get All Users
func (h ReportHandler) HandleGetAllReports(c echo.Context) ([]dashboard.Report, error) {
	supabaseClient, err := functions.CreateSupabaseClient()

	if err != nil {
		fmt.Println(err)
		fmt.Print("Error HandleGetAllReports Line 103")
		return nil, c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	var results []map[string]interface{}

	// Get all users
	err = supabaseClient.DB.From("cake").Select("*").Execute(&results)
	if err != nil {
		fmt.Println(err)
		fmt.Print("Error HandleGetAllReports Line 113")
		return nil, c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Format the results
	var reports []dashboard.Report
	for _, result := range results {

		var incidentDate string

		if result["incident_date"] == nil {
			incidentDate = "No date"
		} else {
			incidentDate = fmt.Sprintf("%s", result["incident_date"])
		}

		report := dashboard.Report{
			IncidentDate:          incidentDate,
			TypeOfReport:          result["typeOfReport"].(string),
			ClientName:            result["clientName"].(string),
			ClientSurname:         result["clientSurname"].(string),
			ClientAddress:         result["clientAddress"].(string),
			RespondingOfficerName: result["responseName"].(string),
			ResponderCallSign:     result["responderCallSign"].(string),
			ResponderArrivalTime:  result["responderArrivalTime"].(string),
			OperatorName:          result["operatorName"].(string),
			OperatorPosition:      result["operatorPosition"].(string),
			Report:                result["report"].(string),
			UserId:                result["user_id"].(string),
		}
		reports = append(reports, report)
	}

	return reports, nil
}

// Dashboard - Users - Table Page
func (h ReportHandler) HandleDashboardReportsTablePage(c echo.Context) error {
	reports, err := h.HandleGetAllReports(c)

	if err != nil {
		fmt.Println(err)
		fmt.Print("Error HandleGetAllUser Line 155")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return render(c, dashboard.DashboardReportsTablePage(reports))
}
