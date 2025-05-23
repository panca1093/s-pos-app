package response

import (
	"github.com/labstack/echo/v4"
)

// Success Response is a utility function to send JSON responses in an Echo framework.
func Success(c echo.Context, code int, data interface{}) error {
	return c.JSON(code, map[string]interface{}{
		"success": true,
		"data":    data,
	})
}

// Error Response is a utility function to send JSON error responses in an Echo framework.
func Error(c echo.Context, code int, msg string) error {
	return c.JSON(code, map[string]interface{}{
		"success": false,
		"error":   msg,
	})
}
