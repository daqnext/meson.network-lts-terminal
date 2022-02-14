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

var DisablePortMap = map[int]struct{}{
	1:    {},
	7:    {},
	9:    {},
	11:   {},
	13:   {},
	15:   {},
	17:   {},
	19:   {},
	20:   {},
	21:   {},
	22:   {},
	23:   {},
	25:   {},
	37:   {},
	42:   {},
	43:   {},
	53:   {},
	77:   {},
	79:   {},
	80:   {},
	87:   {},
	95:   {},
	101:  {},
	102:  {},
	103:  {},
	104:  {},
	109:  {},
	110:  {},
	111:  {},
	113:  {},
	115:  {},
	117:  {},
	119:  {},
	123:  {},
	135:  {},
	139:  {},
	143:  {},
	179:  {},
	389:  {},
	465:  {},
	512:  {},
	513:  {},
	514:  {},
	515:  {},
	526:  {},
	530:  {},
	531:  {},
	532:  {},
	540:  {},
	556:  {},
	563:  {},
	587:  {},
	601:  {},
	636:  {},
	993:  {},
	995:  {},
	2049: {},
	3659: {},
	4045: {},
	6000: {},
	6665: {},
	6666: {},
	6667: {},
	6668: {},
	6669: {},
	6697: {},
}

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
	if echoServer != nil {
		return nil
	}
	http_port, err := configuration.Config.GetInt("port", 8080)
	if err != nil {
		return errors.New("port [int] in config error," + err.Error())
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
	s.Echo.Server.SetKeepAlivesEnabled(false)
	return s.Echo.Start(":" + strconv.Itoa(s.Http_port))
}

func (s *EchoServer) Restart() error {
	s.Shutdown()
	time.Sleep(5 * time.Second)
	s.Echo.Server.SetKeepAlivesEnabled(false)
	return s.Start()
}

func (s *EchoServer) StartTLS(certFile, keyFile interface{}) error {
	basic.Logger.Debugln("https server started on port :" + strconv.Itoa(s.Http_port))
	s.Echo.TLSServer.SetKeepAlivesEnabled(false)
	return s.Echo.StartTLS(":"+strconv.Itoa(s.Http_port), certFile, keyFile)
}

func (s *EchoServer) RestartTls(certFile, keyFile interface{}) error {
	s.Shutdown()
	time.Sleep(5 * time.Second)
	s.Echo.TLSServer.SetKeepAlivesEnabled(false)
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
