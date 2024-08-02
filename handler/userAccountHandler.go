package handler

import (
	"context"
	"fmt"
	"net/http"

	supa "github.com/nedpals/supabase-go"

	"github.com/kerneels94/reports/functions"
	"github.com/kerneels94/reports/view/auth"
	"github.com/labstack/echo/v4"
)

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
}

func AddUserAccount() {
	row := User{
		ID:        user.ID,
		FirstName: c.FormValue("name"),
		LastName:  c.FormValue("surname"),
		Role:      "admin",
	}

	// Add data to the user table
	var results []User
	err = supabaseClient.DB.From("users").Insert(row).Execute(&results)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "An error occurred while creating user"})
	}
}