package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validate *validator.Validate
}

func NewCustomValidator(validator *validator.Validate) echo.Validator {
	return &CustomValidator{validate: validator}
}

func (c *CustomValidator) Validate(i any) error {
	return c.validate.Struct(i)
}
