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
	api := e.Group("/api")
	{
		admin := api.Group("/admin")
		{
			admin.POST("/new", handler.Admin.New)
			admin.POST("/login", handler.Admin.Login)
			admin.GET("/logout", handler.Admin.Logout)
		}
		account := api.Group("/account")
		{
			account.POST("/new", handler.Account.New)
			account.POST("/modifystate", handler.Account.ModifyState)
			account.POST("/info", handler.Account.Info)
			account.GET("/list", handler.Account.List)
			account.POST("/recharge", handler.Account.Recharge)
		}
		gym := api.Group("gym")
		{
			gym.POST("/new", handler.Ground.New)
			gym.POST("/info", handler.Ground.Info)
			gym.GET("/list", handler.Ground.List)
			gym.POST("/state", handler.Ground.ModifyState)
		}
		service := api.Group("/service")
		{
			service.POST("/in", handler.Service.In)
			service.POST("/out", handler.Service.Out)
		}
		bill := api.Group("/bill")
		{
			bill.POST("/info", handler.Bill.Info)
			bill.GET("/list", handler.Bill.List)
			bill.POST("/listbyphone", handler.Bill.ListByPhone)
			bill.POST("/listbygid", handler.Bill.ListByGid)
		}

	}
}

// todo: without filter
