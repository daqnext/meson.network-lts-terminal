package controllers

import (
	"io/ioutil"
	"net/http"
	"path/filepath"

	meson_msg "github.com/daqnext/meson-msg"
	"github.com/daqnext/meson.network-lts-terminal/basic"
	"github.com/labstack/echo/v4"
	"github.com/universe-30/UUtils/path_util"
)

// @Summary      save file command
// @Description  save file from given url
// @Tags         server cmd
// @Accept       json
// @Produce      json
// @Param        SaveFileMsg  body   meson_msg.SaveFileMsg  true "save command object"
// @Param        Signature  header  string  true  "sdfwefwfwfwfsdfwfwf"
// @Success      200  {string}  string  "{"msg": "hello  Razeen"}"
// @Failure      400  {string}  string  "{"msg": "who    are  you"}"
// @Failure      401  {string}  string  "Unauthorized"
// @Router       /api/save [post]
func saveHandler(ctx echo.Context) error {
	var msg meson_msg.SaveFileMsg
	if err := ctx.Bind(&msg); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.String(200, msg.NameHash)
}

// @Summary      delete file command
// @Description  delete file on terminal disk
// @Tags         server cmd
// @Produce      json
// @Param        nameHash  path  string  true  "0dea69026ee1c698"
// @Param        Signature  header  string  true  "sdfwefwfwfwfsdfwfwf"
// @Success      200  {string}  string  "{"msg": "hello  Razeen"}"
// @Failure      400  {string}  string  "{"msg": "who    are  you"}"
// @Failure      401  {string}  string  "Unauthorized"
// @Router       /api/delete/:nameHash [get]
func deleteHandler(ctx echo.Context) error {
	nameHash := ctx.Param("nameHash")
	if len(nameHash) != 16 {
		return echo.NewHTTPError(http.StatusBadRequest, "namehash error")
	}
	return ctx.String(200, nameHash)
}

// @Summary      delete file command
// @Description  delete file on terminal disk
// @Tags         server cmd
// @Produce      json
// @Param        nameHash  path  string  true  "0dea69026ee1c698"
// @Param        Signature  header  string  true  "sdfwefwfwfwfsdfwfwf"
// @Success      200  {string}  string  "{"msg": "hello  Razeen"}"
// @Failure      400  {string}  string  "{"msg": "who    are  you"}"
// @Failure      401  {string}  string  "Unauthorized"
// @Router       /api/delete/:nameHash [get]
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

// @Summary      delete file command
// @Description  delete file on terminal disk
// @Tags         server cmd
// @Produce      json
// @Param        nameHash  path  string  true  "0dea69026ee1c698"
// @Param        Signature  header  string  true  "sdfwefwfwfwfsdfwfwf"
// @Success      200  {string}  string  "{"msg": "hello  Razeen"}"
// @Failure      400  {string}  string  "{"msg": "who    are  you"}"
// @Failure      401  {string}  string  "Unauthorized"
// @Router       /api/delete/:nameHash [get]
func checkLogHandler(ctx echo.Context) error {
	file := ctx.Param("*")
	return ctx.File(filepath.Join(path_util.GetAbsPath("./logs"), file))
}

// @Summary      restart node command
// @Description  restart node command
// @Tags         server cmd
// @Produce      json
// @Param        nameHash  path  string  true  "0dea69026ee1c698"
// @Param        Signature  header  string  true  "sdfwefwfwfwfsdfwfwf"
// @Success      200  {string}  string  "{"msg": "hello  Razeen"}"
// @Failure      400  {string}  string  "error msg"
// @Failure      401  {string}  string  "Unauthorized"
// @Router       /api/restart [get]
func restartHandler(ctx echo.Context) error {
	return ctx.String(200, "")
}

// @Summary      check ScheduleJob running status
// @Description  check ScheduleJob running status
// @Tags         server cmd
// @Produce      json
// @Param        nameHash  path  string  true  "0dea69026ee1c698"
// @Param        Signature  header  string  true  "sdfwefwfwfwfsdfwfwf"
// @Success      200  {string}  string  "{"msg": "hello  Razeen"}"
// @Failure      400  {string}  string  "error msg"
// @Failure      401  {string}  string  "Unauthorized"
// @Router       /api/schedulejobstatus [get]
func scheduleJobStatusHandler(ctx echo.Context) error {
	return ctx.String(200, "")
}

// @Summary      get node status
// @Description  get node status
// @Tags         server cmd
// @Produce      json
// @Param        Signature  header  string  true  "sdfwefwfwfwfsdfwfwf"
// @Success      200  {string}  string  "{"msg": "hello  Razeen"}"
// @Failure      400  {string}  string  "error msg"
// @Failure      401  {string}  string  "Unauthorized"
// @Router       /api/nodestatus [get]
func nodeStatusHandler(ctx echo.Context) error {
	return ctx.String(200, "")
}
