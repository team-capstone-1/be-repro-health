package middleware

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"capstone-project/config"
)

func CreateToken(userId uuid.UUID, role string) (string, error) {
	claims := jwt.MapClaims{}
	// token kedua (payload)
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	// token pertama (header)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// return bersama token ketiga (dengan secret key)
	return token.SignedString([]byte(config.JWT_KEY))
}

func CheckRole(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			if !user.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]any{
					"message":  "unauthorized",
					"response": "Permission Denied: User is not valid",
				})
			}

			claims := user.Claims.(jwt.MapClaims)
			userRole := claims["role"].(string)

			if userRole == role {
				return next(c)
			}

			return c.JSON(http.StatusUnauthorized, map[string]any{
				"message":  "unauthorized",
				"response": "Permission Denied: Only " + role + " roles are allowed to perform this operation.",
			})
		}
	}
}

func ExtractTokenUserId(e echo.Context) uuid.UUID {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["user_id"].(string)
		uuid, _ := uuid.Parse(userId)
		return uuid
	}
	return uuid.Nil
}