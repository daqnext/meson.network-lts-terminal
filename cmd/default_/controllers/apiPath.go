package controllers

import (
	"github.com/daqnext/meson.network-lts-terminal/cmd/default_/controllers/myMiddleware"
	"github.com/daqnext/meson.network-lts-terminal/src/echoServer"
)

func DeclareApi() {
	httpServer := echoServer.GetSingleInstance()

	//file request
	httpServer.GET("*", cacheFileRequestHandler, myMiddleware.ParseFileRequest)

	//forward request
	//echoServer.Any("/api/forward")

	//server cmd
	httpServer.POST("/api/save", saveHandler, myMiddleware.CheckSign)
	httpServer.POST("/api/delete", deleteHandler, myMiddleware.CheckSign)
	httpServer.GET("/api/checklog", listLogFileHandler, myMiddleware.CheckSign)
	httpServer.GET("/api/checklog/*", checkLogHandler, myMiddleware.CheckSign)

	//speed test
	httpServer.GET("/api/pause/:second", pauseHandler, myMiddleware.CheckSign)

	//health and test
	httpServer.GET("/api/health", healthHandler)
	httpServer.GET("/api/test", testHandler)

	//other
	httpServer.GET("/favicon.ico", faviconHandler)
}
