/*
 * Revision History:
 *     Initial: 2018/05/25        Chen Yanchen
 */

package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/JonSnow47/Gymnasium-management-system/GMS/common"
	"github.com/JonSnow47/Gymnasium-management-system/GMS/model"
	"github.com/JonSnow47/Gymnasium-management-system/GMS/util"
)

type billHandler struct{}

var Bill *billHandler

func (*billHandler) Info(c echo.Context) error {
	var req struct {
		Id string `validate:"alphanum,len=24"`
	}

	req.Id = c.FormValue("id")

	if err := c.Validate(&req); err != nil {
		c.Logger().Error("[Validate]", err)
		return c.JSON(http.StatusOK, Resp(common.RespFailed, common.ErrValidate))
	}

	a, err := model.BillService.Info(req.Id)
	if err != nil {
		c.Logger().Error("[Info]", err)
		return c.JSON(http.StatusOK, Resp(common.RespFailed, common.ErrMongo))
	}

	return c.JSON(http.StatusOK, Resp(common.RespSuccess, a))
}

func (*billHandler) ListByPhone(c echo.Context) error {
	phone := c.FormValue("phone")

	if !util.PhoneNum(phone) {
		c.Logger().Error("[Validate]", common.ErrParam)
		return c.JSON(http.StatusOK, Resp(common.RespFailed, common.ErrValidate))
	}

	bills, err := model.BillService.ListByPhone(phone)
	if err != nil {
		c.Logger().Error("[ListByPhone]", err)
		return c.JSON(http.StatusOK, Resp(common.RespFailed, common.ErrMongo))
	}

	return c.JSON(http.StatusOK, Resp(common.RespSuccess, bills))
}

func (*billHandler) ListByGid(c echo.Context) error {
	gid, err := strconv.Atoi(c.FormValue("id"))

	if err != nil {
		c.Logger().Error("[Validate]", err)
		return c.JSON(http.StatusOK, Resp(common.RespFailed, common.ErrValidate))
	}

	a, err := model.BillService.ListByGid(gid)
	if err != nil {
		c.Logger().Error("[ListByPid]", err)
		return c.JSON(http.StatusOK, Resp(common.RespFailed, common.ErrMongo))
	}

	return c.JSON(http.StatusOK, Resp(common.RespSuccess, a))
}

func (*billHandler) List(c echo.Context) error {

	a, err := model.BillService.List()
	if err != nil {
		c.Logger().Error("[List]", err)
		return c.JSON(http.StatusOK, Resp(common.RespFailed, common.ErrMongo))
	}

	return c.JSON(http.StatusOK, Resp(common.RespSuccess, a))
}
