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
	//httpServer.Any("/api/forward")

	//server cmd
	httpServer.POST("/api/save", saveHandler, myMiddleware.CheckSign)
	httpServer.GET("/api/delete/:nameHash", deleteHandler, myMiddleware.CheckSign)
	httpServer.GET("/api/restart", restartHandler, myMiddleware.CheckSign)
	httpServer.GET("/api/schedulejobstatus", scheduleJobStatusHandler, myMiddleware.CheckSign)
	httpServer.GET("/api/nodestatus", nodeStatusHandler, myMiddleware.CheckSign)

	//log
	httpServer.GET("/api/checklog", listLogFileHandler)
	httpServer.GET("/api/checklog/*", checkLogHandler)

	//speed test
	httpServer.GET("/api/pause/:second", pauseHandler, myMiddleware.CheckSign)

	//health and test
	httpServer.GET("/api/health", healthHandler)
	httpServer.GET("/api/test", testHandler)

	//other
	httpServer.GET("/favicon.ico", faviconHandler)
}
