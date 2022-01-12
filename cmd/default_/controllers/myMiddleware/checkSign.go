package myMiddleware

import (
	"log"

	"github.com/labstack/echo/v4"
)

//func RequestToken(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		log.Println("request token left", global.RequestLimiter.GetTokenLeft())
//		allowRequest := global.RequestLimiter.GetRequestToken()
//		defer global.RequestLimiter.ReleaseRequestToken()
//
//		if !allowRequest {
//			//request full
//			log.Println("request full")
//			return nil
//		}
//
//		if err := next(c); err != nil {
//			log.Println("after")
//			return err
//		}
//		log.Println("after")
//		return nil
//	}
//}

func CheckSign(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("befor")

		if err := next(c); err != nil {
			c.Error(err)
		}
		log.Println("after")
		return nil
	}
}
