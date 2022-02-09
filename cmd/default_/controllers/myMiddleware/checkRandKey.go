package myMiddleware

import (
	"errors"
	"strings"

	"github.com/daqnext/meson.network-lts-terminal/src/randomKeyMgr"
	"github.com/labstack/echo/v4"
)

func CheckRandomKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//get random key in path
		v := strings.Split(c.Request().RequestURI, "mesonrandkey=")
		if len(v) != 2 {
			//return c.String(http.StatusUnauthorized, "invalid random key")
			c.Error(errors.New("invalid access key"))
			return nil
		}
		//check random key
		if !randomKeyMgr.GetSingleInstance().CheckRandomKey(v[1]) {
			//return c.String(http.StatusUnauthorized, "invalid random key")
			c.Error(errors.New("invalid access key"))
			return nil
		}

		//fileName := c.QueryParam("filename")
		//fileName = strings.TrimSuffix(fileName, "mesonrandkey="+v[1])
		//log.Println("fileName:", fileName)

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
