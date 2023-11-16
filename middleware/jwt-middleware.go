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

func CreateTokenDoctor(adminID uuid.UUID, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["adminID"] = adminID
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWT_KEY))

}

func ExtractTokenDoctor(c echo.Context) (uuid.UUID, error) {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return uuid.Nil, echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}
	claims := user.Claims.(jwt.MapClaims)
	if claims["admin"] != true {
		return uuid.Nil, echo.NewHTTPError(401, "Unauthorized")
	}

	adminID := claims["adminID"].(string)
	uid, err := uuid.Parse(adminID)
	if err != nil {
		return uuid.Nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	return uid, nil
}
