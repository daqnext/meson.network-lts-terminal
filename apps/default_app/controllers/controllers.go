package controllers

import (
	"github.com/daqnext/meson.network-lts-terminal/apps/default_app/global"
	"github.com/labstack/echo/v4"
	"time"

	_ "github.com/daqnext/meson.network-lts-terminal/apps/default_app/controllers/cacheFileRequest"
	_ "github.com/daqnext/meson.network-lts-terminal/apps/default_app/controllers/health"
	_ "github.com/daqnext/meson.network-lts-terminal/apps/default_app/controllers/serverCmd"
	_ "github.com/daqnext/meson.network-lts-terminal/apps/default_app/controllers/speedTest"
)

func init() {
	if !global.GLOBAL_INIT_FINISHED {
		return
	}

	global.EchoServer.Echo.GET("/favicon.ico", faviconHandler)
}

func faviconHandler(ctx echo.Context) error {
	return ctx.String(200, time.Now().String())
}
