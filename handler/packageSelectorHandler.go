package handler

import (
	"fmt"
	"net/http"

	"github.com/kerneels94/reports/config"
	"github.com/kerneels94/reports/functions"
	"github.com/kerneels94/reports/view/tiers"
	"github.com/labstack/echo/v4"
)

type PackageType struct {
	isLoggedIn  bool
	packageType string
}

type UserWithPackage struct {
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	HasPackage bool   `json:"has_package"`
	ID         string `json:"id"`
	IsLoggedIn bool   `json:"is_logged_in"`
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
			functions.HtmxRedirect(c, "/dashboard")
			break
		}

		if oneHundred != "" {
			results[1].packageType = "oneHundred"
			supabaseClient.DB.From("tiers").Insert(results[1]).Execute(&results)
			functions.HtmxRedirect(c, "/dashboard")
			break
		}

		if twoHundred != "" {
			results[1].packageType = "twoHundred"
			supabaseClient.DB.From("tiers").Insert(results[1]).Execute(&results)
			functions.HtmxRedirect(c, "/dashboard")
			break
		}

		if unlimited != "" {
			results[1].packageType = "unlimited"
			supabaseClient.DB.From("tiers").Insert(results[1]).Execute(&results)
			functions.HtmxRedirect(c, "/dashboard")
			break
		}

	}

	return nil
}

// Render package selector page
func (h PackageType) RenderPackagePage(c echo.Context) error {
	return render(c, tiers.PackageSelected())
}

func (h PackageType) GetUserPackage(c echo.Context) error {
	// Check if the user has a package
	supabaseClient, err := functions.CreateSupabaseClient()

	// Get user id
	userId, err := config.GetUserIdFromCookie(c.Request(), c.Echo().AcquireContext(), supabaseClient)

	// Get object of user
	var results UserWithPackage
	err = supabaseClient.DB.From("users").Select("*").Single().Eq("id", userId).Execute(&results)

	// Check if there is an error
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// check if results is empty if so the user has no package and redirect to the tiers
	if results.HasPackage {
		fmt.Println("Has package")
		// functions.HtmxRedirect(c, "/dashboard")
	} else {
		fmt.Println("No package found")
		// fmt.Println(results)
		// functions.HtmxRedirect(c, "/select-package")
	}

	// return JSON
	return c.JSON(http.StatusOK, results)
}
