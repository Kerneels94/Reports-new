package handler

// todo - globalize the API_URL and API_KEY
// todo - create a function to check if the API_URL and API_KEY are set
// todo - create a function to initialize the client
// todo - create a function to fetch data
// todo - create a function to unmarshal data

import (
	"fmt"
	"net/http"
	"os"

	supa "github.com/nedpals/supabase-go"

	"github.com/labstack/echo/v4"
)

type NotesHandler struct{}

func (h NotesHandler) HandleNotes(c echo.Context) error {
	// API_URL
	API_URL := os.Getenv("API_URL")
	if API_URL == "" {
		fmt.Println("API_URL is not set")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "API_URL is not set"})
	}

	// API_KEY
	API_KEY := os.Getenv("API_KEY")
	if API_KEY == "" {
		fmt.Println("API_KEY is not set")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "API_KEY is not set"})
	}

	supabaseClient := supa.CreateClient(API_URL, API_KEY)

	var results []map[string]interface{}
	err := supabaseClient.DB.From("notes").Select("*").Execute(&results)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "An error occurred"})
	}

	print(results)
	return c.JSON(http.StatusOK, results)
}
