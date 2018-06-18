/*
 * Revision History:
 *     Initial: 2018/05/25        Chen Yanchen
 */

package handler

import (
	"github.com/labstack/echo"
	"net/http"

	"github.com/JonSnow007/Gymnasium-management-system/GMS/common"
	"github.com/JonSnow007/Gymnasium-management-system/GMS/model"
	"github.com/JonSnow007/Gymnasium-management-system/GMS/util"
	"gopkg.in/mgo.v2"
)

type billHandler struct{}

var Bill *billHandler

func (*billHandler) Info(c echo.Context) error {
	var req struct {
		Id string `validate:"alphanum,len=24"`
	}

	//req.Id = c.FormValue("id")
	if err := c.Bind(&req); err != nil {
		c.Logger().Error("[Bind]", err)
		return c.JSON(http.StatusOK, Resp(common.ErrParam))
	}

	if err := c.Validate(&req); err != nil {
		c.Logger().Error("[Validate]", err)
		return c.JSON(http.StatusOK, Resp(common.ErrValidate))
	}

	a, err := model.BillService.Info(req.Id)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusOK, Resp(common.ErrNotFound))
		}
		c.Logger().Error("[Info]", err)
		return c.JSON(http.StatusOK, Resp(common.ErrMongoDB))
	}

	return c.JSON(http.StatusOK, Resp(common.RespSuccess, a))
}

func (*billHandler) ListByPhone(c echo.Context) error {
	var req struct {
		Phone string `json:"phone" validate:"required,len=11"`
	}

	//req.Phone = c.FormValue("phone")
	if err := c.Bind(&req); err != nil {
		c.Logger().Error("[Bind]", err)
		return c.JSON(http.StatusOK, Resp(common.ErrParam))
	}

	if !util.PhoneNum(req.Phone) {
		c.Logger().Error("[Validate]", common.ErrParam)
		return c.JSON(http.StatusOK, Resp(common.ErrValidate))
	}

	bills, err := model.BillService.ListByPhone(req.Phone)
	if err != nil {
		c.Logger().Error("[ListByPhone]", err)
		return c.JSON(http.StatusOK, Resp(common.ErrMongoDB))
	}

	return c.JSON(http.StatusOK, Resp(common.RespSuccess, bills))
}

func (*billHandler) ListByGid(c echo.Context) error {
	var req struct {
		Id int `json:"id" validate:"require,numeric"`
	}

	if err := c.Bind(&req); err != nil {
		c.Logger().Error("[Bind]", err)
		return c.JSON(http.StatusOK, Resp(common.ErrValidate))
	}

	a, err := model.BillService.ListByGid(req.Id)
	if err != nil {
		c.Logger().Error("[ListByPid]", err)
		return c.JSON(http.StatusOK, Resp(common.ErrMongoDB))
	}

	return c.JSON(http.StatusOK, Resp(common.RespSuccess, a))
}

func (*billHandler) List(c echo.Context) error {
	a, err := model.BillService.List()
	if err != nil {
		c.Logger().Error("[List]", err)
		return c.JSON(http.StatusOK, Resp(common.ErrMongoDB))
	}

	return c.JSON(http.StatusOK, Resp(common.RespSuccess, a))
}

func (*billHandler) Total(c echo.Context) error {
	total, err := model.BillService.Total()
	if err != nil {
		return c.JSON(http.StatusOK, Resp(common.ErrMongoDB, nil))
	}

	recorded, err := model.AccountService.Recorded()
	if err != nil {
		return c.JSON(http.StatusOK, Resp(common.ErrMongoDB, nil))
	}
	return c.JSON(http.StatusOK, Resp(common.RespSuccess, map[string]int{
		"total":    total,
		"recorded": recorded,
	}))
}
