package server

import (
	"net/http"

	"github.com/ardafirdausr/posjoo-server/internal/entity"
	"github.com/labstack/echo/v4"
)

type CustomHTTPErrorHandler struct {
	debug  bool
	logger echo.Logger
}

func NewCustomHTTPErrorHandler(debug bool, logger echo.Logger) *CustomHTTPErrorHandler {
	customErrHandler := new(CustomHTTPErrorHandler)
	customErrHandler.debug = debug
	customErrHandler.logger = logger
	return customErrHandler
}

func (che CustomHTTPErrorHandler) Handler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if !ok {
		he = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	if ev, ok := err.(entity.ErrValidation); ok {
		he.Code = http.StatusBadRequest
		he.Message = ev.Message
	}

	if ent, ok := err.(entity.ErrInvalidData); ok {
		he.Code = http.StatusBadRequest
		he.Message = ent.Message
	}

	if ent, ok := err.(entity.ErrUnauthorize); ok {
		he.Code = http.StatusUnauthorized
		he.Message = ent.Message
	}

	if ent, ok := err.(entity.ErrNotFound); ok {
		he.Code = http.StatusNotFound
		he.Message = ent.Message
	}

	if he.Message == "" {
		he.Message = http.StatusText(he.Code)
	}

	code := he.Code
	payload := he.Message
	if m, ok := he.Message.(string); ok {
		if che.debug {
			payload = echo.Map{"message": m, "error": err.Error()}
		} else {
			payload = echo.Map{"message": m}
		}
	}

	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(he.Code)
		} else {
			err = c.JSON(code, payload)
		}
		if err != nil {
			che.logger.Error(err)
		}
	}
}
