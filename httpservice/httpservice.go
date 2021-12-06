package httpservice

import (
	"github.com/daqnext/meson.network-lts-http-server"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

func RunHttpService() {
	hs := httpserver.New()

	// if is debug mode
	if true {
		hs.SetLogLevel_DEBUG()
	}

	//hs.SetLogFile("./log")
	hs.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "-> ${time_rfc3339_nano} | ${method} | ${remote_ip} | ${host} | ${uri} | ${status} | ${latency_human}\n",
	}))

	//cors
	hs.Use(middleware.CORS())

	//api
	hs.GET("/api/testapi/test", func(ctx httpserver.Context) error {
		return ctx.JSON(http.StatusOK, nil)
	})

	//hs.GET("/*", func(ctx httpserver.Context) error {
	//	filePath:=ctx.Path()
	//	host:=ctx.Request().Host
	//	ip:=ctx.RealIP()
	//
	//
	//})

	//if err := hs.StartTLS(":8443", "server.crt", "server.key"); err != http.ErrServerClosed {
	//	log.Fatalln(err)
	//}

	if err := hs.Start(":8080"); err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}
