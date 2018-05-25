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
	e.GET("/account/modifystate", handler.Account.ModifyState)
	e.GET("/account/info", handler.Account.Info)
	e.GET("/account/list", handler.Account.List)
	e.GET("/account/recharge", handler.Account.Recharge)

	e.GET("/ground/new", handler.Ground.New)
	e.GET("/ground/info", handler.Ground.Info)
	e.GET("/ground/list", handler.Ground.List)
	e.GET("/ground/state", handler.Ground.ModifyState)

	//////////////////未测试
	e.GET("/bill/info", handler.Bill.Info)
	e.GET("/bill/list", handler.Bill.List)

	//e.GET("/in", handler.In)
	//e.GET("/out", handler.Out)

	e.GET("/test", handler.Test)
}
