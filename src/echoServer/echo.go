package echoServer

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/daqnext/MesonTerminalEchoServer"
	"github.com/daqnext/meson.network-lts-terminal/basic"
	"github.com/daqnext/meson.network-lts-terminal/configuration"
	"github.com/daqnext/meson.network-lts-terminal/tools"
	"github.com/labstack/echo/v4/middleware"
	"github.com/universe-30/EchoMiddleware"
	"github.com/universe-30/EchoMiddleware/tool"
)

type EchoServer struct {
	*MesonTerminalEchoServer.HttpServer
	Http_port int
}

var echoServer *EchoServer

func GetSingleInstance() *EchoServer {
	return echoServer
}

/*
http_port
http_static_rel_folder
*/
func Init() error {
	http_port, err := configuration.Config.GetInt("http_port", 8080)
	if err != nil {
		return errors.New("http_port [int] in config error," + err.Error())
	}

	s := &EchoServer{
		MesonTerminalEchoServer.New(),
		http_port,
	}

	//cros
	s.Use(middleware.CORS())
	//logger
	s.Use(EchoMiddleware.LoggerWithConfig(EchoMiddleware.LoggerConfig{
		Logger:            basic.Logger,
		RecordFailRequest: true,
	}))
	//recover and panicHandler
	s.Use(EchoMiddleware.RecoverWithConfig(EchoMiddleware.RecoverConfig{
		OnPanic: tools.PanicHandler,
	}))

	s.UseJsoniter()

	echoServer = s

	return nil
}

//use jsoniter
func (s *EchoServer) UseJsoniter() {
	s.JSONSerializer = tool.NewJsoniter()
}

func (s *EchoServer) Start() error {
	basic.Logger.Debugln("http server started on port :" + strconv.Itoa(s.Http_port))
	return s.Echo.Start(":" + strconv.Itoa(s.Http_port))
}

func (s *EchoServer) Restart() error {
	s.Shutdown()
	time.Sleep(5 * time.Second)
	return s.Start()
}

func (s *EchoServer) StartTLS(certFile, keyFile interface{}) error {
	basic.Logger.Debugln("https server started on port :" + strconv.Itoa(s.Http_port))
	return s.Echo.StartTLS(":"+strconv.Itoa(s.Http_port), certFile, keyFile)
}

func (s *EchoServer) RestartTls(certFile, keyFile interface{}) error {
	s.Shutdown()
	time.Sleep(5 * time.Second)
	return s.StartTLS(certFile, keyFile)
}

func (s *EchoServer) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer func() {
		s.Echo.Listener = nil
		s.Echo.Server = new(http.Server)
		s.Echo.Server.SetKeepAlivesEnabled(false)
		s.Echo.Server.Handler = s.Echo

		s.Echo.TLSListener = nil
		s.Echo.TLSServer = new(http.Server)
		s.Echo.TLSServer.SetKeepAlivesEnabled(false)
		s.Echo.TLSServer.Handler = s.Echo
	}()

	if err := s.Echo.Shutdown(ctx); err != nil {
		s.CloseServer()
		basic.Logger.Errorln("http server shutdown error:", err)
	} else {
		basic.Logger.Debugln("shutdown processed successfully")
	}
}
