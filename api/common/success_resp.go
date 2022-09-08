package common

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ResSuccessWithData(ctx echo.Context, data interface{}) error {
	if data == nil {
		data = []interface{}{}
	}

	return ctx.JSON(
		http.StatusOK,
		struct {
			Code    int         `json:"code"`
			Message string      `json:"message"`
			Data    interface{} `json:"data"`
		}{
			Code:    http.StatusOK,
			Message: "request successful",
			Data:    data,
		})
}

func ResSuccessCreated(ctx echo.Context, msg string) error {
	return ctx.JSON(
		http.StatusCreated,
		struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{Code: http.StatusCreated, Message: msg},
	)
}

func ResSuccessOK(ctx echo.Context, msg string) error {
	return ctx.JSON(
		http.StatusOK,
		struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{Code: http.StatusOK, Message: msg},
	)
}
