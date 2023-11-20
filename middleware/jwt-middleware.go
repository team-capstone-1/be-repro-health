package middleware

import (
	"errors"
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
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"message":  "unauthorized",
					"response": "Permission Denied: User is not valid",
				})
			}

			claims := user.Claims.(jwt.MapClaims)
			userRole := claims["role"].(string)

			if userRole == role {
				return next(c)
			}

			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
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

func ExtractTokenDoctor(c echo.Context) (uuid.UUID, error) {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return uuid.Nil, echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}
	claims := user.Claims.(jwt.MapClaims)
	if claims["role"] != "doctor" {
		return uuid.Nil, echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	userId := claims["user_id"].(string)
	uid, err := uuid.Parse(userId)
	if err != nil {
		return uuid.Nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	return uid, nil
}

func ExtractTokenDoctorId(e echo.Context) (uuid.UUID, error) {
	doctor := e.Get("doctor")
	if doctor == nil {
		// Handle case where "doctor" is nil, perhaps return an error or take appropriate action
		return uuid.Nil, errors.New("doctor token is nil")
	}

	doctorToken, ok := doctor.(*jwt.Token)
	if !ok || !doctorToken.Valid {
		// Handle case where "doctor" is not a valid *jwt.Token
		return uuid.Nil, errors.New("invalid doctor token")
	}

	claims, ok := doctorToken.Claims.(jwt.MapClaims)
	if !ok {
		// Handle case where claims cannot be extracted
		return uuid.Nil, errors.New("unable to extract claims from doctor token")
	}

	doctorID, ok := claims["doctor_id"].(string)
	if !ok {
		// Handle case where "doctor_id" is not present or not a string
		return uuid.Nil, errors.New("doctor_id not present or not a string")
	}
	
	parsedUUID, err := uuid.Parse(doctorID)
	if err != nil {
		// Handle case where parsing UUID fails
		return uuid.Nil, err
	}
	return parsedUUID, nil
}

