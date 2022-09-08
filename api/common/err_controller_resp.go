package common

import "github.com/labstack/echo/v4"

func ResErrorController(ctx echo.Context, code int, msg string) error {
	return ctx.JSON(code, &ResponseForm1{Code: code, Message: msg})
}
