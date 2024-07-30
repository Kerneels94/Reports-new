package main

import (
	// "context"
	"fmt"

	"github.com/joho/godotenv"

	"github.com/kerneels94/reports/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	app := echo.New()

	// app.Use(withUser)
	
	userHandler := handler.UserHandler{}
	app.GET("/user", userHandler.HandleUserShow)
	
	loginHandler := handler.LoginHandler{}
	app.GET("/login", loginHandler.HandleUserLogin)

	signUpHandler := handler.SignUpHandler{}
	app.GET("/sign-up", signUpHandler.HandleSignUp)
	app.POST("/sign-up", signUpHandler.HandleUserSignUp)

	noteHandler := handler.NotesHandler{}
	app.GET("/notes", noteHandler.HandleNotes)

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
