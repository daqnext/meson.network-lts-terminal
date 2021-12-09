package speedTest

import (
	"github.com/daqnext/meson.network-lts-terminal/apps/default_app/global"
	"github.com/labstack/echo/v4"
	"time"
)

func init() {
	if !global.GLOBAL_INIT_FINISHED {
		return
	}

	global.EchoServer.Echo.POST("/api/pause", pauseHandler)
}

func pauseHandler(ctx echo.Context) error {
	global.EchoServer.Echo.SetPauseSeconds(4)
	return ctx.String(200, time.Now().String())
}
