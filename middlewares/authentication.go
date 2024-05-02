package middlewares

import (
	"net/http"

	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/user"
	"github.com/labstack/echo/v4"
)

func Authentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			verifyToken, err := user.VerifyToken(c)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error":   "Unauthenticated",
					"message": err.Error(),
				})
			}
			c.Set("userData", verifyToken)
			return next(c)
		}
	}
}
