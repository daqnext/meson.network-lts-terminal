package controllers

import (
	"time"

	"github.com/labstack/echo/v4"
)

func pauseHandler(ctx echo.Context) error {
	//global.HttpServer.SetPauseSeconds(4)
	return ctx.String(200, time.Now().String())
}
