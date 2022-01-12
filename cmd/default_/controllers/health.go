package controllers

import (
	"time"

	"github.com/labstack/echo/v4"
)

func healthHandler(ctx echo.Context) error {
	return ctx.String(200, time.Now().String())
}

func testHandler(ctx echo.Context) error {
	return ctx.String(200, time.Now().String())
}
