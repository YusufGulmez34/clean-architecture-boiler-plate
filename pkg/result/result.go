package result

import "github.com/labstack/echo/v4"

func NewSuccessResult(c echo.Context, statusCode int, message string) error {
	return Response(c, statusCode, echo.Map{"success": true, "message": message})
}

func NewErrorResult(c echo.Context, statusCode int, message string, err error) error {
	return Response(c, statusCode, echo.Map{"success": false, "message": message + ": " + err.Error()})
}

func NewDataResult(c echo.Context, statusCode int, message string, data any) error {
	return Response(c, statusCode, echo.Map{"success": true, "message": message, "data": data})
}

func Response(c echo.Context, statusCode int, data any) error {
	c.Response().Header().Set("content-type", "application/json")
	return c.JSON(statusCode, data)
}
