/*
 * Revision History:
 *     Initial: 2018/05/21        Chen Yanchen
 */

package router

import (
	"github.com/labstack/echo"

	"github.com/JonSnow47/Gymnasium-management-system/GMS/handler"
)

func Init(e *echo.Echo) {
	e.GET("/home", handler.Home)

	e.POST("/admin/new", handler.Admin.New)
	e.POST("/admin/login", handler.Admin.Login)

	e.GET("/account/new", handler.Account.New)
	e.GET("/account/modifyphone", handler.Account.ModifyPhone)
	e.GET("/account/modifystate", handler.Account.ModifyState)
	e.GET("/account/io", handler.Account.InOut)
	e.GET("/account/info",handler.Account.Info)

	e.GET("/test",handler.Account.Test)
}
