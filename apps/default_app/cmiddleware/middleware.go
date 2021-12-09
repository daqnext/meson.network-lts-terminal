package cmiddleware

import (
	"github.com/labstack/echo/v4"
	"log"
)

func CheckRandomKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("befor")
		//check random key
		c.Set("randomKey", "1234321")

		if err := next(c); err != nil {
			c.Error(err)
		}
		log.Println("after")
		return nil
	}

}

func CheckSign2(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("befor")
		//check random key

		if err := next(c); err != nil {
			c.Error(err)
		}
		log.Println("after")
		return nil
	}
}

func CheckSign() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log.Println("befor")
			//check random key

			if err := next(c); err != nil {
				c.Error(err)
			}
			log.Println("after")
			return nil
		}

	}
}
