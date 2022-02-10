package controllers

import (
	"github.com/labstack/echo/v4"
)

//func pathCacheFileRequestHandler(ctx echo.Context) error {
//	//check bindName fileName
//
//	//get bindname
//	rPath := ctx.Param("*")
//	s := strings.SplitN(rPath, "/", 2)
//
//	if len(s) < 1 {
//		return ctx.String(200, "url Error:"+ctx.Request().RequestURI)
//	}
//	bindName := s[0]
//	if bindName == "" {
//		return ctx.String(200, "url Error:"+ctx.Request().RequestURI)
//	}
//	fileName := "index.html"
//	if len(s) > 1 && s[1] != "" {
//		fileName = s[1]
//	}
//
//	basic.Logger.Debugln(bindName, fileName)
//
//	return ctx.String(200, ctx.Request().RequestURI)
//}
//
//func domainCacheFileRequestHandler(ctx echo.Context) error {
//	//check bindName fileName
//	//time.Sleep(10 * time.Second)
//
//	uri := ctx.Request().RequestURI
//	if uri == "/" {
//		uri = "/index.html"
//	}
//
//	log.Println("main")
//	randomKey := ctx.Get("randomKey")
//	log.Println(randomKey.(string))
//
//	newAdd := "https://local.shoppynext.com:10443/api/cdn" + uri
//	return ctx.Redirect(302, newAdd)
//
//	return ctx.String(200, uri)
//}

// @Summary      get cached file
// @Description  get cached file
// @Tags         public
// @Success      200  {string}  string  "{"msg": "hello  Razeen"}"
// @Failure      400  {string}  string  "error msg"
// @Failure      401  {string}  string  "Unauthorized"
// @Router       /anyPath [get]
func cacheFileRequestHandler(ctx echo.Context) error {
	bindName := ctx.Get("bindName").(string)
	fileName := ctx.Get("fileName").(string)
	fileHash := ctx.Get("fileHash").(string)

	return ctx.String(200, "bindName:"+bindName+"\nfileName:"+fileName+"\nfileHash:"+fileHash)
}
