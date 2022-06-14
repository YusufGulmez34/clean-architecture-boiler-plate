package middlewares

import "github.com/labstack/echo/v4"

type IMiddleware interface {
	AuthControl(next echo.HandlerFunc) echo.HandlerFunc
}
