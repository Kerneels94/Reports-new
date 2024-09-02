package functions

import (
	"net/http"

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
