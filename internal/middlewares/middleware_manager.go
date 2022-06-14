package middlewares

import (
	"boiler-plate/internal/storage"
	"boiler-plate/pkg/constants"
	"boiler-plate/pkg/global"
	"boiler-plate/pkg/result"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type MiddlewareManager struct {
	storage storage.IStorage
}

func NewMiddlewareManager(storage storage.IStorage) IMiddleware {
	return &MiddlewareManager{storage: storage}
}

func (m *MiddlewareManager) AuthControl(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return result.NewErrorResult(c, http.StatusUnauthorized, constants.AuthError, errors.New("Token not faund"))
		}

		tokenString = strings.Split(tokenString, " ")[1]
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(global.TokenSecretKey), nil
		})

		if err != nil {
			return result.NewErrorResult(c, http.StatusUnauthorized, constants.AuthError, err)
		}

		claims := token.Claims.(jwt.MapClaims)
		if _, err = m.storage.TokenDetails().GetByUUID(claims["uuid"].(string)); err != nil {
			return result.NewErrorResult(c, http.StatusUnauthorized, constants.AuthError, err)
		}

		if _, err = m.storage.Users().FindByID(uint(claims["user_id"].(float64))); err != nil {
			return result.NewErrorResult(c, http.StatusUnauthorized, constants.AuthError, err)
		}

		return next(c)
	}
}
