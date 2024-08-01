package functions

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	supa "github.com/nedpals/supabase-go"
)

func checkApiCreds() (string, string) {
	// Get the API_URL and API_KEY from the environment variables
	API_URL := os.Getenv("API_URL")
	API_KEY := os.Getenv("API_KEY")

	// Check if the API_URL and API_KEY is not set
	if API_URL == "" || API_KEY == "" {
		fmt.Println("API_URL is not set or API_KEY is not set")
		return "", ""
	}

	return API_URL, API_KEY
}

func CreateSupabaseClient() (*supa.Client, error) {
	API_URL, API_KEY := checkApiCreds()

	if API_URL == "" || API_KEY == "" {
		return nil, fmt.Errorf("API_URL is not set or API_KEY is not set")
	}

	return supa.CreateClient(API_URL, API_KEY), nil
}

func HtmxRedirect(c echo.Context, url string) {
	c.Response().Header().Set("HX-Redirect", url)
}
