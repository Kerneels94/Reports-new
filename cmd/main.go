package main

import (
	"context"

	"github.com/kerneels94/reports/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	userHandler := handler.UserHandler{}

	app.Use(withUser)

	app.GET("/user", userHandler.HandleUserShow)

	app.Start(":3000")

}

// Middleware
func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("user", "test@gmail.com")
		ctx := context.WithValue(c.Request().Context(), "user", "test@gmail.com")
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
