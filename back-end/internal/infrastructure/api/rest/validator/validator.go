package validator

import (
	"github.com/labstack/echo/v4"
   "github.com/go-playground/validator/v10"
)


type wrapper struct {
	validator *validator.Validate
}

func New()  echo.Validator {
	return &wrapper {
		validator: validator.New(),
	}
}

func (w *wrapper) Validate(i interface{}) error {
    return w.validator.Struct(i)
}