package cacheFileRequest

import (
	"github.com/daqnext/meson.network-lts-terminal/apps/default_app/cmiddleware"
	"github.com/daqnext/meson.network-lts-terminal/apps/default_app/global"
	"github.com/daqnext/meson.network-lts-terminal/cli"
	"github.com/labstack/echo/v4"
	"log"
	"strings"
	"time"
)

func init() {
	if !global.GLOBAL_INIT_FINISHED {
		return
	}

	global.EchoServer.Echo.GET("/api/cdn/*", pathCacheFileRequestHandler, cmiddleware.CheckSign2)
	global.EchoServer.Echo.GET("*", domainCacheFileRequestHandler, cmiddleware.CheckRandomKey)
}

func pathCacheFileRequestHandler(ctx echo.Context) error {
	//check bindName fileName

	//get bindname
	rPath := ctx.Param("*")
	s := strings.SplitN(rPath, "/", 2)

	if len(s) < 1 {
		return ctx.String(200, "url Error:"+ctx.Request().RequestURI)
	}
	bindName := s[0]
	if bindName == "" {
		return ctx.String(200, "url Error:"+ctx.Request().RequestURI)
	}
	fileName := "index.html"
	if len(s) > 1 && s[1] != "" {
		fileName = s[1]
	}

	cli.LocalLogger.Debugln(bindName, fileName)

	return ctx.String(200, ctx.Request().RequestURI)
}

func domainCacheFileRequestHandler(ctx echo.Context) error {
	//check bindName fileName
	time.Sleep(12 * time.Second)

	uri := ctx.Request().RequestURI
	if uri == "/" {
		uri = "/index.html"
	}

	log.Println("main")
	randomKey := ctx.Get("randomKey")
	log.Println(randomKey.(string))

	return ctx.String(200, uri)
}
