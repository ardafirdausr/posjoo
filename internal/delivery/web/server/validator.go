package server

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/ardafirdausr/posjoo-server/internal/entity"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator(validator *validator.Validate) *CustomValidator {
	validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	customValidator := new(CustomValidator)
	customValidator.validator = validator
	return customValidator
}

func (v *CustomValidator) Validate(i interface{}) error {
	err := v.validator.Struct(i)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		verr := entity.ErrValidation{
			Message: "Invalid format data",
			Err:     err,
		}
		if len(validationErrors) > 0 {
			validationError := validationErrors[0]
			validationParam := validationError.Param()
			validationField := validationError.Field()

			switch validationError.Tag() {
			case "required":
				verr.Message = fmt.Sprintf("%s is required", validationField)
			case "min":
				verr.Message = fmt.Sprintf("Min value of %s is %s", validationField, validationParam)
			case "max":
				verr.Message = fmt.Sprintf("Max value of %s is %s", validationField, validationParam)
			case "eqfield":
				verr.Message = fmt.Sprintf("Value of %s must be equal as %s", validationField, validationParam)
			}
		}

		return verr
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
