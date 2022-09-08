package common

import (
	"belajar-mongodb/business"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ResBusinessError(ctx echo.Context, err error) error {
	switch err {
	default:
		return errBusinessResJSON(ctx, http.StatusInternalServerError, err.Error())

	case business.ErrDataNotValidate:
		return errBusinessResJSON(ctx, http.StatusBadRequest, err.Error())

	case business.ErrNotFound:
		return errBusinessResJSON(ctx, http.StatusNotFound, err.Error())

	case business.ErrConflict:
		return errBusinessResJSON(ctx, http.StatusConflict, err.Error())

	}
}

func errBusinessResJSON(ctx echo.Context, code int, message string) error {
	return ctx.JSON(code, &ResponseForm1{Code: code, Message: message})
}
