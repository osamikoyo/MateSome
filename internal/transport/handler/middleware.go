package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/osamikoyo/matesome/internal/data/models"
	"net/http"
	"os"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		claims := &models.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return os.Getenv("JWT_KEY"), nil
		})

		if !token.Valid || err != nil {
			return c.JSON(http.StatusUnauthorized, "invalid token")
		}

		c.Set("email", claims.Email)
		return next(c)
	}
}
