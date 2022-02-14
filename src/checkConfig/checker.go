package checkConfig

import (
	"errors"
	"fmt"
	"strings"

	meson_msg "github.com/daqnext/meson-msg"
	"github.com/daqnext/meson.network-lts-terminal/basic"
	"github.com/daqnext/meson.network-lts-terminal/src/destMgr"
	"github.com/daqnext/meson.network-lts-terminal/src/echoServer"
	"github.com/daqnext/meson.network-lts-terminal/src/requestUtil"
	"github.com/go-playground/validator/v10"
)

func CheckLogLevel(newLevel string) error {
	allLevel := map[string]struct{}{
		"TRAC":  {},
		"TRACE": {},
		"DEBU":  {},
		"DEBUG": {},
		"INFO":  {},
		"WARN":  {},
		"ERRO":  {},
		"ERROR": {},
		"PANI":  {},
		"PANIC": {},
		"FATA":  {},
		"FATAL": {},
	}
	level := strings.ToUpper(newLevel)
	_, exist := allLevel[level]
	if !exist {
		return errors.New("input level error. should be TRAC DEBU INFO WARN ERRO PANI or PATA")
	}
	return nil
}

func CheckPort(port int) error {
	if port <= 0 {
		return errors.New("input port must greater than 0")
	}

	_, exist := echoServer.DisablePortMap[port]
	if exist {
		return fmt.Errorf("port [%d] is forbidden", port)
	}

	//get location from server
	location := "US"
	location = strings.ToLower(location)
	if location == "china" || location == "cn" {

	}

	return nil
}

func CheckToken(token string) error {
	destMgr.Init()
	err := destMgr.GetSingleInstance().SearchAvailableDest()
	if err != nil {
		return errors.New("can not find server, please set correct server host")
	}

	url := destMgr.GetSingleInstance().GetDestUrl("/api/tokenvalidator")
	resp, err := requestUtil.Post(url, nil, meson_msg.ValidateToken{Token: token}, 15, "")
	if err != nil {
		return errors.New("network err")
	}

	_ = resp

	//if not pass
	if true {
		return errors.New("token invalid")
	}

	return nil
}

func CheckDest(newDest string) error {
	validate := validator.New()

	errs := validate.Var(newDest, "hostname_rfc1123")

	if errs != nil {
		basic.Logger.Debugln(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		return errors.New("dest name not correct")
	}

	return nil
}
