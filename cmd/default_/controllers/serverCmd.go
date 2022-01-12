package controllers

import (
	"time"

	"github.com/labstack/echo/v4"
)

func saveHandler(ctx echo.Context) error {
	return ctx.String(200, time.Now().String())
}

func deleteHandler(ctx echo.Context) error {
	return ctx.String(200, time.Now().String())
}

func checkLogHandler(ctx echo.Context) error {
	return ctx.String(200, time.Now().String())
}
