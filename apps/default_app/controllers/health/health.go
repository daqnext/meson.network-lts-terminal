package health

import (
	"github.com/daqnext/meson.network-lts-terminal/apps/default_app/global"
	"github.com/labstack/echo/v4"
	"time"
)

func init() {
	if !global.GLOBAL_INIT_FINISHED {
		return
	}

	global.EchoServer.Echo.GET("/api/health", healthHandler)
	global.EchoServer.Echo.GET("/api/test", testHandler)
}

func healthHandler(ctx echo.Context) error {
	return ctx.String(200, time.Now().String())
}

func testHandler(ctx echo.Context) error {
	return ctx.String(200, time.Now().String())
}
