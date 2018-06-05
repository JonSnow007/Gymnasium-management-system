/*
 * Revision History:
 *     Initial: 2018/05/21        Chen Yanchen
 */

package router

import (
	"github.com/labstack/echo"

	"github.com/JonSnow007/Gymnasium-management-system/GMS/handler"
)

func Init(e *echo.Echo) {
	e.GET("/api/home", handler.Service.Home)
	e.GET("/api/service/in", handler.Service.In)
	e.GET("/api/service/out", handler.Service.Out)

	e.POST("/api/admin/new", handler.Admin.New)
	e.POST("/api/admin/login", handler.Admin.Login)

	e.GET("/api/account/new", handler.Account.New)
	e.GET("/api/account/modifystate", handler.Account.ModifyState)
	e.GET("/api/account/info", handler.Account.Info)
	e.GET("/api/account/list", handler.Account.List)
	e.GET("/api/account/recharge", handler.Account.Recharge)

	e.GET("/api/gym/new", handler.Ground.New)
	e.GET("/api/gym/info", handler.Ground.Info)
	e.GET("/api/gym/list", handler.Ground.List)
	e.GET("/api/gym/state", handler.Ground.ModifyState)

	e.GET("/api/bill/info", handler.Bill.Info)
	e.GET("/api/bill/list", handler.Bill.List)
	e.GET("/api/bill/listbyphone", handler.Bill.ListByPhone)
	e.GET("/api/bill/listbygid", handler.Bill.ListByGid)
}
