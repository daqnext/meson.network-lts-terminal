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
	1:    struct{}{},
	7:    struct{}{},
	9:    struct{}{},
	11:   struct{}{},
	13:   struct{}{},
	15:   struct{}{},
	17:   struct{}{},
	19:   struct{}{},
	20:   struct{}{},
	21:   struct{}{},
	22:   struct{}{},
	23:   struct{}{},
	25:   struct{}{},
	37:   struct{}{},
	42:   struct{}{},
	43:   struct{}{},
	53:   struct{}{},
	77:   struct{}{},
	79:   struct{}{},
	80:   struct{}{},
	87:   struct{}{},
	95:   struct{}{},
	101:  struct{}{},
	102:  struct{}{},
	103:  struct{}{},
	104:  struct{}{},
	109:  struct{}{},
	110:  struct{}{},
	111:  struct{}{},
	113:  struct{}{},
	115:  struct{}{},
	117:  struct{}{},
	119:  struct{}{},
	123:  struct{}{},
	135:  struct{}{},
	139:  struct{}{},
	143:  struct{}{},
	179:  struct{}{},
	389:  struct{}{},
	465:  struct{}{},
	512:  struct{}{},
	513:  struct{}{},
	514:  struct{}{},
	515:  struct{}{},
	526:  struct{}{},
	530:  struct{}{},
	531:  struct{}{},
	532:  struct{}{},
	540:  struct{}{},
	556:  struct{}{},
	563:  struct{}{},
	587:  struct{}{},
	601:  struct{}{},
	636:  struct{}{},
	993:  struct{}{},
	995:  struct{}{},
	2049: struct{}{},
	3659: struct{}{},
	4045: struct{}{},
	6000: struct{}{},
	6665: struct{}{},
	6666: struct{}{},
	6667: struct{}{},
	6668: struct{}{},
	6669: struct{}{},
	6697: struct{}{},
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
