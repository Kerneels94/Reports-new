package main

import (
	// "context"

	"fmt"

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
	app.GET("/notes", noteHandler.HandleNotes)

	dashboardHandler := handler.DashboardHandler{}
	app.GET("/dashboard", dashboardHandler.HandleDashboard)
	reportHandler := handler.CreateReportHandler{}
	app.GET("/reports", reportHandler.HandleCreateReport)
	app.POST("/api/logout", dashboardHandler.HandleLogout)

	app.Start(":3000")

}

// Middleware
// func withUser(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		c.Set("user", "test@gmail.com")
// 		ctx := context.WithValue(c.Request().Context(), "user", "test@gmail.com")
// 		c.SetRequest(c.Request().WithContext(ctx))
// 		return next(c)
// 	}
// }
