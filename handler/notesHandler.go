package handler

// todo - globalize the API_URL and API_KEY
// todo - create a function to check if the API_URL and API_KEY are set
// todo - create a function to initialize the client
// todo - create a function to fetch data
// todo - create a function to unmarshal data

import (
	"fmt"
	"net/http"

	"github.com/kerneels94/reports/functions"

	"github.com/labstack/echo/v4"
)

type NotesHandler struct{}

func (h NotesHandler) HandleNotes(c echo.Context) error {
	supabaseClient, err := functions.CreateSupabaseClient()

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	var results []map[string]interface{}
	err = supabaseClient.DB.From("notes").Select("*").Execute(&results)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	print(results)
	return c.JSON(http.StatusOK, results)
}
