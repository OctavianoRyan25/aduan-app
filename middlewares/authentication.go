package middlewares

import (
	"net/http"

	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/constants"
	"github.com/OctavianoRyan25/lapor-lingkungan-hidup/modules/user"
	"github.com/labstack/echo/v4"
)

func Authentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Verifikasi token JWT
			claims, err := user.VerifyToken(c)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error":   "Unauthenticated",
					"message": err.Error(),
				})
			}

			// Dapatkan userID dari klaim token
			userIDFloat64, ok := claims["id"].(float64)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error":   "Unauthenticated",
					"message": "Invalid user ID in token",
				})
			}
			userID := uint(userIDFloat64)

			// Dapatkan userID dari klaim token
			role, ok := claims["role"].(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error":   "Unauthenticated",
					"message": "Invalid user ID in token",
				})
			}

			// Setel Role dalam konteks Echo sesuai dengan peran yang diminta
			c.Set("role", role)

			// Setel userID dalam konteks Echo sesuai dengan peran yang diminta
			c.Set("user_id", userID)

			// Lanjutkan ke handler berikutnya
			return next(c)
		}
	}
}

func AuthenticationAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Verifikasi token JWT
			claims, err := user.VerifyToken(c)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error":      constants.ErrUnauthenticated,
					"error_code": constants.ErrCodeUnauthenticated,
					"message":    err.Error(),
				})
			}

			// Dapatkan userID dari klaim token
			userIDFloat64, ok := claims["id"].(float64)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error":      constants.ErrUnauthenticated,
					"error_code": constants.ErrCodeUnauthenticated,
					"message":    "Invalid user ID in token",
				})
			}
			userID := uint(userIDFloat64)

			// Setel userID dalam konteks Echo
			c.Set("admin_id", userID)

			// Lanjutkan ke handler berikutnya
			return next(c)
		}
	}
}
