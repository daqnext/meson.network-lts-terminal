package controllers

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	"github.com/daqnext/meson.network-lts-terminal/basic"
	"github.com/labstack/echo/v4"
	"github.com/universe-30/UUtils/path_util"
)

func saveHandler(ctx echo.Context) error {
	return ctx.String(200, time.Now().String())
}

func deleteHandler(ctx echo.Context) error {
	return ctx.String(200, time.Now().String())
}

func listLogFileHandler(ctx echo.Context) error {
	logFiles := []byte{}
	//all logs
	path := path_util.GetAbsPath("./logs/all")
	rd, err := ioutil.ReadDir(path)
	if err != nil {
		basic.Logger.Errorln("read ./logs/all fail", "err:", err, "path:", path)
		return echo.NewHTTPError(http.StatusNoContent, "list all log files err", err)
	}
	for _, fi := range rd {
		if !fi.IsDir() {
			name := "<a href=" + "/api/checklog/all/" + fi.Name() + ">" + "logs/all/" + fi.Name() + "</a><br/>"
			logFiles = append(logFiles, []byte(name)...)
		}
	}
	logFiles = append(logFiles, []byte("<hr/>")...)

	//error logs
	path = path_util.GetAbsPath("./logs/error")
	rde, err := ioutil.ReadDir(path)
	if err != nil {
		basic.Logger.Errorln("read ./logs/error fail", "err:", err, "path:", path)
		return echo.NewHTTPError(http.StatusNoContent, "list error log files err", err)
	}
	for _, fi := range rde {
		if !fi.IsDir() {
			name := "<a href=" + "/api/checklog/error/" + fi.Name() + ">" + "logs/error/" + fi.Name() + "</a><br/>"
			logFiles = append(logFiles, []byte(name)...)
		}
	}

	return ctx.Blob(http.StatusOK, "text/html; charset=utf-8", logFiles)
}

func checkLogHandler(ctx echo.Context) error {
	file := ctx.Param("*")
	return ctx.File(filepath.Join(path_util.GetAbsPath("./logs"), file))
}
