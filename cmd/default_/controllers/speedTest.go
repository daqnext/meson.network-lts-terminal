package controllers

import (
	"net/http"
	"strconv"

	"github.com/daqnext/meson.network-lts-terminal/src/echoServer"
	"github.com/labstack/echo/v4"
)

// @Summary      pause file transfer
// @Description  pause file transfer for several seconds, and do speed test
// @Tags         server cmd
// @Produce      json
// @Param        second  path  string  true  "4"
// @Param        Signature  header  string  true  "sdfwefwfwfwfsdfwfwf"
// @Success      200  {string}  string  "{"msg": "hello  Razeen"}"
// @Failure      400  {string}  string  "error msg"
// @Failure      401  {string}  string  "Unauthorized"
// @Router      /api/pause/:second [get]
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
