package controllers

import (
	"github.com/daqnext/meson.network-lts-terminal/cmd/default_/controllers/myMiddleware"
	"github.com/daqnext/meson.network-lts-terminal/components/echoServer"
)

func RunApi() {
	echoServer := echoServer.GetSingleInstance()
	//file request
	echoServer.GET("/api/cdn/*", pathCacheFileRequestHandler, myMiddleware.CheckRandomKey)
	echoServer.GET("*", domainCacheFileRequestHandler, myMiddleware.CheckRandomKey)

	//health and test
	echoServer.GET("/api/health", healthHandler)
	echoServer.GET("/api/test", testHandler)

	//server cmd
	echoServer.POST("/api/save", saveHandler, myMiddleware.CheckSign)
	echoServer.POST("/api/delete", deleteHandler, myMiddleware.CheckSign)
	echoServer.POST("/api/checklog", checkLogHandler, myMiddleware.CheckSign)

	//speed test
	echoServer.POST("/api/pause", pauseHandler, myMiddleware.CheckSign)

	//other
	echoServer.GET("/favicon.ico", faviconHandler)
}
