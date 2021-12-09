package components

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"

	localLog "github.com/daqnext/LocalLog/log"
	fj "github.com/daqnext/fastjson"

	echoserver "github.com/daqnext/meson.network-lts-http-server"
	"github.com/daqnext/utils/path_util"
)

type EchoServer struct {
	Echo         *echoserver.HttpServer
	Http_port    int
	localLogger  *localLog.LocalLog
	certFilePath string
	keyFilePath  string
}

/*
http_port
http_static_rel_folder
*/
func InitEchoServer(localLogger_ *localLog.LocalLog, ConfigJson *fj.FastJson) (*EchoServer, error) {

	http_port, err := ConfigJson.GetInt("http_port")
	if err != nil {
		return nil, errors.New("http_port [int] in config.json not defined," + err.Error())
	}

	certFile, err := ConfigJson.GetString("cert_path")
	if err != nil {
		return nil, errors.New("cert_path [string] in config.json not defined," + err.Error())
	}

	keyFile, err := ConfigJson.GetString("key_path")
	if err != nil {
		return nil, errors.New("key_path [string] in config.json not defined," + err.Error())
	}

	es := &EchoServer{
		echoserver.New(),
		http_port,
		localLogger_,
		path_util.GetAbsPath(certFile),
		path_util.GetAbsPath(keyFile),
	}

	//set locallogger
	//es.Echo.Use(NewEchoLogger(localLogger_))
	return es, nil
}

func (s *EchoServer) Start() error {
	s.localLogger.Infoln("http server started on port :" + strconv.Itoa(s.Http_port))

	return s.Echo.StartTLS(":"+strconv.Itoa(s.Http_port), s.certFilePath, s.keyFilePath)
}

func (s *EchoServer) Close() {
	s.Echo.Close()
}

func (s *EchoServer) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.Echo.Shutdown(ctx); err != nil {
		s.localLogger.Errorln("http server shutdown error:", err)
		s.Close()
	}
}

func neglectErrors(status int) bool {
	if status == 404 || status == 405 {
		return true
	} else {
		return false
	}
}

func NewEchoLogger(l *localLog.LocalLog) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			if err != nil {
				c.Error(err)
				//don't log the err if it is 404 requests
				//too many noise requests
				if !neglectErrors(c.Response().Status) || l.Level >= localLog.LLEVEL_DEBUG {
					l.WithFields(localLog.Fields{
						"request":     c.Request().RequestURI,
						"method":      c.Request().Method,
						"remote":      c.Request().RemoteAddr,
						"status":      c.Response().Status,
						"text_status": http.StatusText(c.Response().Status),
						"took":        time.Since(start),
						"request_id":  c.Request().Header.Get("X-Request-Id"),
					}).Errorln("request error:" + err.Error())
				}
			} else {
				//info log ,only for loglevels : debug or trace
				if l.Level >= localLog.LLEVEL_DEBUG {
					lentry := l.WithFields(localLog.Fields{
						"request":     c.Request().RequestURI,
						"method":      c.Request().Method,
						"remote":      c.Request().RemoteAddr,
						"status":      c.Response().Status,
						"text_status": http.StatusText(c.Response().Status),
						"took":        time.Since(start),
						"request_id":  c.Request().Header.Get("X-Request-Id"),
					})
					if l.Level == localLog.LLEVEL_DEBUG {
						lentry.Infoln("request success")
					} else {
						lentry.Traceln("request success")
					}
				}
			}
			return nil
		}
	}
}
