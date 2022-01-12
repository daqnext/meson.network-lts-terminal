package controllers

import (
	"github.com/labstack/echo/v4"
)

func faviconHandler(ctx echo.Context) error {
	return ctx.String(200, "")
}
