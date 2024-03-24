package middlewares

import (
	"net/http"
	"task/models"
	"task/responses"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Path() == "/register" || c.Path() == "/login" {
			return next(c)
		}
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return responses.ErrorResponse(c, http.StatusUnauthorized, responses.UNAUTHORIZED)
		}
		token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return models.JWTSecret, nil
		})
		if err != nil {
			return responses.ErrorResponse(c, http.StatusUnauthorized, responses.INVALID_TOKEN)
		}
		if !token.Valid {
			return responses.ErrorResponse(c, http.StatusUnauthorized, responses.INVALID_TOKEN)
		}
		claims, ok := token.Claims.(*models.CustomClaims)
		if !ok {
			return responses.ErrorResponse(c, http.StatusUnauthorized, responses.INVALID_TOKEN)
		}

		// Add the authenticated user information to the Echo context
		c.Set("user_id", claims.UserID)

		return next(c)
	}
}
