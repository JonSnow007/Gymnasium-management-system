/*
 * Revision History:
 *     Initial: 2018/05/19        Chen Yanchen
 */

package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"gopkg.in/go-playground/validator.v9"

	"github.com/JonSnow007/Gymnasium-management-system/GMS/router"
)

func main() {
	initEcho()
}

func initEcho() {
	e := echo.New()
	// 设置静态资源文件路径
	e.Static("/", "static")
	// 中间件
	// e.Use(middleware.session)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	// 跨站配置
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
	// 	AllowCredentials: true,
	//   }))

	e.HideBanner = true
	e.Logger.SetLevel(log.Lvl(1))
	// 定制 Validator, 基于validator.v9
	e.Validator = &CustomValidator{validator: validator.New()}

	router.Init(e)

	e.Logger.Fatal(e.Start(":8080"))
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func init() {
	var first, last string
	fmt.Println("Who am I?")
	fmt.Scanf("%s %s", &first, &last)
	if first != "Jon" || last != "Snow" {
		fmt.Println("Made in earth by humans.")
		os.Exit(0)
	}
}
