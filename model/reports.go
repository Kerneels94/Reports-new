package model

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