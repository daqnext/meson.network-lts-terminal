package myMiddleware

import (
	"errors"

	"github.com/daqnext/meson.network-lts-terminal/cmd/default_/manager/randomKeyMgr"
	"github.com/labstack/echo/v4"
)

func CheckRandomKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//get random key in path

		//check random key
		if !randomKeyMgr.GetSingleInstance().CheckRandomKey("123") {
			c.Error(errors.New("access key error"))
			return nil
		}

		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}

}
