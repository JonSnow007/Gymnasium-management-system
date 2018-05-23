/*
 * Revision History:
 *     Initial: 2018/05/22        Chen Yanchen
 */

package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/JonSnow47/Gymnasium-management-system/GMS/common"
	"github.com/JonSnow47/Gymnasium-management-system/GMS/model"
	"strconv"
)

type accountHandler struct{}

var Account *accountHandler

func (*accountHandler) New(c echo.Context) error {
	var (
		err error
		req struct {
			Name  string `json:"name" validate:"required,min=3,max=10"`
			Phone string `json:"phone" validate:"required,numeric,len=11"`
		}
	)

	if err = c.Bind(&req); err != nil {
		c.Logger().Info("[Parameter]", err)
		return c.JSON(http.StatusOK, RespData(common.RespFailed, common.ErrParam))
	}

	if err = c.Validate(&req); err != nil {
		c.Logger().Info("[Validate]", err)
		return c.JSON(http.StatusOK, RespData(common.RespFailed, common.ErrValidate))
	}

	id, err := model.AccountService.New(req.Name, req.Phone)
	if err != nil {
		c.Logger().Error("[New account]", err)
		return c.JSON(http.StatusOK, RespData(common.RespFailed, common.ErrMongo))
	}

	return c.JSON(http.StatusOK, RespId(common.RespSuccess, id))
}

func (*accountHandler) ModifyPhone(c echo.Context) error {
	var req struct {
		Old string `json:"phone" validate:"required,numeric,len=11"`
		New string `json:"phone" validate:"required,numeric,len=11"`
	}

	req.Old = c.FormValue("old")
	req.New = c.FormValue("new")

	if err := c.Validate(&req); err != nil {
		c.Logger().Info("[Validate]", err)
		return c.JSON(http.StatusOK, RespData(common.RespFailed, common.ErrValidate))
	}

	if err := model.AccountService.ModifyPhone(req.Old, req.New); err != nil {
		c.Logger().Error("[ModifyPhone]", err)
		return c.JSON(http.StatusOK, RespData(common.RespFailed, common.ErrMongo))
	}

	return c.JSON(http.StatusOK, Resp(common.RespSuccess))
}

func (*accountHandler) ModifyState(c echo.Context) error {
	phone := c.FormValue("phone")

	err := model.AccountService.ModifyState(phone)
	if err != nil {
		c.Logger().Error("[ModifyState]", err)
		return c.JSON(http.StatusOK, RespData(common.RespFailed, common.ErrMongo))
	}

	return c.JSON(http.StatusOK, Resp(common.RespSuccess))
}

// 改变进出场状态
func (*accountHandler) InOut(c echo.Context) error {
	var req struct {
		Phone string `validate:"numeric,len=11"`
	}

	req.Phone = c.FormValue("phone")

	err := model.AccountService.InOut(req.Phone)
	if err != nil {
		c.Logger().Error("[ModifyState]", err)
		return c.JSON(http.StatusOK, RespData(common.RespFailed, common.ErrMongo))
	}

	return c.JSON(http.StatusOK, Resp(common.RespSuccess))
}

// 信息查询
func (*accountHandler) Info(c echo.Context) error {
	var req struct {
		Phone string `validate:"numeric,len=11"`
	}

	req.Phone = c.FormValue("phone")

	a, err := model.AccountService.Info(req.Phone)
	if err != nil {
		c.Logger().Error("[ModifyState]", err)
		return c.JSON(http.StatusOK, RespData(common.RespFailed, common.ErrMongo))
	}

	return c.JSON(http.StatusOK, RespData(common.RespSuccess, a))
}

// 充值与支付
func (*accountHandler) Deal(c echo.Context) error {
	var req struct {
		Phone string `validate:"numeric,len=11"`
		Sum   int    `validate:"numeric,"`
	}

	req.Phone = c.FormValue("phone")
	req.Sum, _ = strconv.Atoi(c.FormValue("sum"))

	err := model.AccountService.Deal(req.Phone, req.Sum)
	if err != nil {
		if err.Error() == "Lack of balance" {
			c.Logger().Error("[Deal]", err)
			return c.JSON(http.StatusOK, RespData(common.RespFailed, common.ErrBalance))
		} else {
			c.Logger().Error("[Deal]", err)
			return c.JSON(http.StatusOK, RespData(common.RespFailed, common.ErrDeal))
		}
	}
	return c.JSON(http.StatusOK, Resp(common.RespSuccess))
}

func (*accountHandler) Test(c echo.Context) error {
	model.Test("13633309095")
	return nil
}
