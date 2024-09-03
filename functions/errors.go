package functions

import (
	"net/http"

	"github.com/kerneels94/reports/handler"
	viewError "github.com/kerneels94/reports/view/error"

	"github.com/labstack/echo/v4"
)

// JsonBadReqError - 400
func JsonBadReqError(c echo.Context, message string) error {
	return c.JSON(http.StatusInternalServerError, map[string]string{"error": message})
}

// InternalServerError - 500
func JsonInternalServerError(c echo.Context, message string) error {
	return c.JSON(http.StatusInternalServerError, map[string]string{"error": message})
}

// NotFound - 404
func JsonNotFound(c echo.Context) error {
	return c.JSON(http.StatusNotFound, map[string]string{"error": "Not found"})
}

// Unauthorized - 401
func JsonUnauthorized(c echo.Context) error {
	return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
}

func DisplayUnauthorizedPage(c echo.Context) error {
	return handler.Render(c, viewError.UnauthorizedPage())
}
