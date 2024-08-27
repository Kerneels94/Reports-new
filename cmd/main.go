package main

import (
	// "context"

	"fmt"

	"github.com/kerneels94/reports/config"

	"github.com/joho/godotenv"
	"github.com/kerneels94/reports/handler"
	"github.com/labstack/echo/v4"
)

/*
Main function
*/
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	app := echo.New()

	// app.Static("./css/main.css", "css")

	// app.Use(withUser)

	mainPageHandler := handler.MainPageHandler{}
	app.GET("/", mainPageHandler.HandleShowMainPage)

	loginHandler := handler.LoginHandler{}
	app.GET("/login", loginHandler.HandleUserLogin)
	app.POST("/api/login", loginHandler.HandleUserLoginLogin)

	signUpHandler := handler.SignUpHandler{}
	app.GET("/sign-up", signUpHandler.HandleSignUp)
	app.POST("/api/sign-up", signUpHandler.HandleUserSignUp)

	noteHandler := handler.NotesHandler{}
	app.GET("/notes", noteHandler.HandleNotes, userHasValidSessionMiddleWare)

	// Dashboard
	dashboardHandler := handler.DashboardHandler{}
	app.GET("/dashboard", dashboardHandler.HandleDashboard, userHasValidSessionMiddleWare)

	// Dashboard users
	// app.GET("/dashboard/users", dashboardHandler.HandleUsers)
	app.GET("/dashboard/users", dashboardHandler.HandleDashboardUsersTablePage, userHasValidSessionMiddleWare) // Display users in table
	// app.GET("/api/dashboard/users", dashboardHandler.HandleGetAllUser) // Get users
	app.POST("/api/dashboard/users", dashboardHandler.HandleAddUser, userHasValidSessionMiddleWare) // Add user

	reportHandler := handler.ReportHandler{}
	app.GET("/reports", reportHandler.HandleShowCreateReportForm, userHasValidSessionMiddleWare)
	app.POST("/api/dashboard/reports", reportHandler.HandleCreateReport, userHasValidSessionMiddleWare)

	// logout
	app.POST("/api/logout", dashboardHandler.HandleLogout, userHasValidSessionMiddleWare)

	app.Start(":3000")
}

// Middleware
func userHasValidSessionMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check if the session is valid
		isValid := config.IsCookieValid(c.Request(), c)

		if !isValid {
			return c.JSON(401, map[string]string{"error": "Unauthorized"})
		}

		return next(c)
	}
}
