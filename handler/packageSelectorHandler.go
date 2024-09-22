package handler

import (
	"fmt"
	"net/http"

	"github.com/kerneels94/reports/functions"
	"github.com/labstack/echo/v4"
)

type PackageType struct {
	isLoggedIn  bool
	packageType string
}

func (h PackageType) HandleSelectPackage(c echo.Context) error {

	supabaseClient, err := functions.CreateSupabaseClient()

	if err != nil {
		fmt.Println("Error connecting:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	var results []PackageType

	supabaseClient.DB.From("users").Select("isLoggedIn").Execute(&results)

	switch results[0].isLoggedIn {
	case true:
		functions.HtmxRedirect(c, "/dashboard")
		return c.JSON(http.StatusOK, "User is logged in")
	case false:
		free := c.FormValue("free")
		oneHundred := c.FormValue("100")
		twoHundred := c.FormValue("200")
		unlimited := c.FormValue("unlimited")

		if free == "" || oneHundred == "" || twoHundred == "" || unlimited == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Please select a package"})
		}

		// This will be a little redundent
		if free != "" {
			results[1].packageType = "free"
			supabaseClient.DB.From("tiers").Insert(results[1]).Execute(&results)
		}

		if oneHundred != "" {
			results[1].packageType = "oneHundred"
			supabaseClient.DB.From("tiers").Insert(results[1]).Execute(&results)
		}

		if twoHundred != "" {
			results[1].packageType = "twoHundred"
			supabaseClient.DB.From("tiers").Insert(results[1]).Execute(&results)
		}

		if unlimited != "" {
			results[1].packageType = "unlimited"
			supabaseClient.DB.From("tiers").Insert(results[1]).Execute(&results)
		}

		functions.HtmxRedirect(c, "/tiers")
		return c.JSON(http.StatusOK, "User is logged out")
	}

	return nil
}
