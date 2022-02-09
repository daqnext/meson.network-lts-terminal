package controllers

import (
	"net/http"
	"strconv"

	"github.com/daqnext/meson.network-lts-terminal/src/echoServer"
	"github.com/labstack/echo/v4"
)

func pauseHandler(ctx echo.Context) error {
	value := ctx.Param("second")
	pauseSecond, err := strconv.Atoi(value)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "pause time error:"+err.Error())
	}

	pauseTime := 4
	if pauseSecond > 0 && pauseSecond < 10 {
		pauseTime = pauseSecond
	}

	echoServer.GetSingleInstance().SetPauseSeconds(int64(pauseTime))

	return ctx.JSON(200, nil)
}
