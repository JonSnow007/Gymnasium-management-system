/*
 * Revision History:
 *     Initial: 2018/05/19        Chen Yanchen
 */

package main

import (
	"github.com/labstack/gommon/log"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"

	"github.com/JonSnow47/Gymnasium-management-system/GMS/router"
)

func main() {
	initEcho()
}

//func init() {
//	log.SetFlags(log.Ldate | log.Lmicroseconds) // 设置 log 格式为毫秒级
//	db.ConnectMongo()                           // 尝试连接 MongoDB
//}

func initEcho() {
	e := echo.New()
	// 设置静态资源文件路径
	e.Static("/", "static")
	// 中间件
	// e.Use(middleware.session)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.SetLevel(log.Lvl(1))
	// 定制 Validator, 基于validator.v9
	e.Validator = &CustomValidator{validator: validator.New()}

	router.Init(e)

	e.Logger.Fatal(e.Start(":9999"))
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
