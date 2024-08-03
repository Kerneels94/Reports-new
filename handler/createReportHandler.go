package handler

import (
	"fmt"
	"net/http"

	"github.com/kerneels94/reports/functions"
	"github.com/labstack/echo/v4"
)

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

type CreateReportHandler struct{}

/*
	Handler for creating a report
*/

func (h CreateReportHandler) HandleCreateReport(c echo.Context) error {
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
