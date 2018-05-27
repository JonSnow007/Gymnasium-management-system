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
	e.GET("/home", handler.Service.Home)

	e.POST("/admin/new", handler.Admin.New)
	e.POST("/admin/login", handler.Admin.Login)

	e.GET("/account/new", handler.Account.New)
	e.GET("/account/modifystate", handler.Account.ModifyState)
	e.GET("/account/info", handler.Account.Info)
	e.GET("/account/list", handler.Account.List)
	e.GET("/account/recharge", handler.Account.Recharge)

	e.GET("/gym/new", handler.Ground.New)
	e.GET("/gym/info", handler.Ground.Info)
	e.GET("/gym/list", handler.Ground.List)
	e.GET("/gym/state", handler.Ground.ModifyState)

	e.GET("/bill/info", handler.Bill.Info)
	e.GET("/bill/list", handler.Bill.List)
	e.GET("/bill/listbyphone", handler.Bill.ListByPhone)
	e.GET("/bill/listbygid", handler.Bill.ListByGid)

	e.GET("/service/in", handler.Service.In)
	e.GET("/service/out", handler.Service.Out)
}
