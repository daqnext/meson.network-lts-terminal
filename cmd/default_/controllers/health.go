package controllers

import (
	"time"

	"github.com/labstack/echo/v4"
)

// @Summary      health check
// @Description  health check
// @Tags         public
// @Produce      plain
// @Success      200  {string}  string  "UTC time string"
// @Router       /api/health [get]
func healthHandler(ctx echo.Context) error {
	return ctx.String(200, time.Now().UTC().String())
}

// @Summary      test
// @Description  test api
// @Tags         public
// @Produce      plain
// @Success      200  {string}  string  "UTC time string"
// @Router       /api/test [get]
func testHandler(ctx echo.Context) error {
	return ctx.String(200, time.Now().UTC().String())
}
