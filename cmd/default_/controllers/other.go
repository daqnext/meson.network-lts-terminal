package controllers

import (
	"github.com/labstack/echo/v4"
)

// @Summary      handle favicon request
// @Description  handle favicon request
// @Tags         public
// @Produce      plain
// @Success      200  {string}  string  "empty string"
// @Router      /favicon [get]
func faviconHandler(ctx echo.Context) error {
	return ctx.String(200, "")
}
