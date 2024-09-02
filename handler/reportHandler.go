package handler

import (
	"fmt"
	"net/http"
	"strings"

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
		return functions.JsonInternalServerError(c, err.Error())
	}

	userId, err := config.GetUserIdFromCookie(c.Request(), c.Echo().AcquireContext(), supabaseClient)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error from line 47 handleCreateReport handler")
		return functions.JsonInternalServerError(c, err.Error())
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
	err = supabaseClient.DB.From("reports").Insert(query).Execute(&results)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error from line 71 handleCreateReport handler")
		return functions.JsonBadReqError(c, err.Error())
	}

	return c.JSON(http.StatusOK, "Report created")

}

// Function to display create report form
func (h ReportHandler) HandleShowCreateReportForm(c echo.Context) error {
	return render(c, dashboard.CreateReportForm())
}

// Dashboard - Users - Get all reports
func (h ReportHandler) HandleGetAllReports(c echo.Context) ([]dashboard.Report, error) {
	supabaseClient, err := functions.CreateSupabaseClient()

	if err != nil {
		fmt.Println(err)
		fmt.Print("Error HandleGetAllReports Line 106")
		return nil, functions.JsonInternalServerError(c, err.Error())
	}

	var results []ReportsData

	// Get all reports
	err = supabaseClient.DB.From("reports").Select("*").Execute(&results)
	if err != nil {
		fmt.Println(err)
		fmt.Print("Error HandleGetAllReports Line 116")
		return nil, functions.JsonInternalServerError(c, err.Error())
	}

	var reports []dashboard.Report

	// Format the results
	for _, result := range results {
		reports = append(reports, formatReportData(result))
	}

	return reports, nil
}

// Dashboard - Users - Table Page
func (h ReportHandler) HandleDashboardReportsTablePage(c echo.Context) error {
	reports, err := h.HandleGetAllReports(c)

	if err != nil {
		fmt.Println(err)
		fmt.Print("Error HandleGetAllUser Line 155")
		return functions.JsonInternalServerError(c, err.Error())
	}

	return render(c, dashboard.DashboardReportsTablePage(reports))
}

/*
Helper functions that format the report data
*/
func formatReportData(reportData ReportsData) dashboard.Report {
	report := dashboard.Report{
		IncidentDate:          trimReportDataContent(reportData.IncidentDate, "No date"),
		TypeOfReport:          trimReportDataContent(reportData.TypeOfReport, "No type"),
		ClientName:            trimReportDataContent(reportData.ClientName, "No name"),
		ClientSurname:         trimReportDataContent(reportData.ClientSurname, "No surname"),
		ClientAddress:         trimReportDataContent(reportData.ClientAddress, "No address"),
		RespondingOfficerName: trimReportDataContent(reportData.RespondingOfficerName, "No responder"),
		ResponderCallSign:     trimReportDataContent(reportData.ResponderCallSign, "No call sign"),
		ResponderArrivalTime:  trimReportDataContent(reportData.ResponderArrivalTime, "No time"),
		OperatorName:          trimReportDataContent(reportData.OperatorName, "No operator"),
		OperatorPosition:      trimReportDataContent(reportData.OperatorPosition, "No position"),
		Report:                trimReportDataContent(reportData.Report, "No report"),
		UserId:                trimReportDataContent(reportData.UserId, "No user id"),
	}

	return report
}

/*
Helper function to trim the report data content
*/
func trimReportDataContent(reportContent string, noContentMessage string) string {
	if strings.TrimSpace(reportContent) == "" {
		return noContentMessage
	}
	return fmt.Sprintf("%s", reportContent)
}
