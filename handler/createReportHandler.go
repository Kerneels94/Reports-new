package handler

import (
	"fmt"
	"net/http"

	"github.com/kerneels94/reports/functions"
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
}

type CreateReportHandler struct{}

/*
	Handler for creating a report
*/

func (h CreateReportHandler) HandleCreateReport(c echo.Context) error {
	supabaseClient, err := functions.CreateSupabaseClient()

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error from line 37 handleCreateReport handler")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()}) // return 500
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
	operatorPosition := c.FormValue("operatorPosition")
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
		OperatorPosition:      operatorPosition,
		Report:                report,
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
