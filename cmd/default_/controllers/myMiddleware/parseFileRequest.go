package myMiddleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/daqnext/meson.network-lts-terminal/src/randomKeyMgr"
	"github.com/labstack/echo/v4"
)

const randomKeyMark = "mesonrandkey"

// parseFileRequest check randomKey and get bindName fileName
func ParseFileRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//get random key in path
		v := strings.Split(c.Request().RequestURI, randomKeyMark)
		randKey := ""
		if len(v) == 1 {
			cookie, err := c.Cookie(randomKeyMark)
			if err != nil {
				c.Error(errors.New("invalid access key"))
				return nil
			}
			randKey = cookie.Value
		} else if len(v) == 2 {
			randKey = v[1]
		} else {
			//return c.String(http.StatusUnauthorized, "invalid random key")
			c.Error(errors.New("invalid access key"))
			return nil
		}
		//check random key
		if !randomKeyMgr.GetSingleInstance().CheckRandomKey(randKey) {
			//return c.String(http.StatusUnauthorized, "invalid random key")
			c.Error(errors.New("invalid access key"))
			return nil
		}

		//set cookie
		cookie := new(http.Cookie)
		cookie.Name = randomKeyMark
		cookie.Value, _ = randomKeyMgr.GetSingleInstance().GetRandomKey()
		c.SetCookie(cookie)

		//get fileName
		fileName := v[0][1:]
		//check fileName legal

		c.Set("fileName", fileName)

		//get bindName
		bindName := parseBindName(c.Request().Host)
		c.Set("bindName", bindName)

		//get fileHash
		fileHash := bindName + fileName //to hash
		c.Set("fileHash", fileHash)

		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}

func parseBindName(hostName string) string {
	return ""
}
