package middleware

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"

	"capstone-project/config"
)

func CreateToken(userId uuid.UUID, role, name string, is_web bool) (string, error) {
	claims := jwt.MapClaims{}
	// token kedua (payload)
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["name"] = name
	claims["role"] = role
	if is_web {
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	}
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

// MiddlewareMock is a mock implementation for middleware testing
type MiddlewareMock struct {
	mock.Mock
}

// CheckRole mocks the behavior of the CheckRole middleware
func (m *MiddlewareMock) CheckRole(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			args := m.Called(role, c)
			return args.Error(0)
		}
	}
}

// ExtractTokenUserId mocks the behavior of the ExtractTokenUserId middleware
func (m *MiddlewareMock) ExtractTokenUserId(c echo.Context) uuid.UUID {
	args := m.Called(c)
	return args.Get(0).(uuid.UUID)
}
