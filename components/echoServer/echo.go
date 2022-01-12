package echoServer

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/daqnext/MesonTerminalEchoServer"
	"github.com/daqnext/meson.network-lts-terminal/basic"
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
var once sync.Once

func GetSingleInstance() *EchoServer {
	once.Do(func() {
		var err error
		echoServer, err = newEchoServer()
		if err != nil {
			basic.Logger.Fatalln(err)
		}
	})
	return echoServer
}

/*
http_port
http_static_rel_folder
*/
func newEchoServer() (*EchoServer, error) {
	http_port, err := basic.Config.GetInt("http_port", 8080)
	if err != nil {
		return nil, errors.New("http_port [int] in config error," + err.Error())
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

	return s, nil
}

//use jsoniter
func (s *EchoServer) UseJsoniter() {
	s.JSONSerializer = tool.NewJsoniter()
}

func (s *EchoServer) Start() error {
	basic.Logger.Infoln("http server started on port :" + strconv.Itoa(s.Http_port))
	return s.Echo.Start(":" + strconv.Itoa(s.Http_port))
}

func (s *EchoServer) StartTLS(certFile, keyFile interface{}) {
	go func() {
		basic.Logger.Infoln("http server started on port :" + strconv.Itoa(s.Http_port))
		err := s.Echo.StartTLS(":"+strconv.Itoa(s.Http_port), certFile, keyFile)
		if err != nil {
			basic.Logger.Errorln(err)
		}
	}()
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
