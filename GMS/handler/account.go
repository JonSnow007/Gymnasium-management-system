/*
 * Revision History:
 *     Initial: 2018/05/22        Chen Yanchen
 */

package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"

	"github.com/JonSnow007/Gymnasium-management-system/GMS/common"
	"github.com/JonSnow007/Gymnasium-management-system/GMS/model"
	"github.com/JonSnow007/Gymnasium-management-system/GMS/util"
)

type accountHandler struct{}

var Account *accountHandler

// 新建用户
func (*accountHandler) New(c echo.Context) error {
	var (
		err error
		req struct {
			Name  string `validate:"required,min=1,max=10"`
			Phone string `validate:"required,numeric,len=11"`
		}
	)

	req.Name = c.FormValue("name")
	req.Phone = c.FormValue("phone")

	if err = c.Validate(&req); err != nil {
		c.Logger().Error("[Validate]", err)
		return c.JSON(http.StatusOK, Resp(common.ErrValidate))
	}

	err = model.AccountService.New(req.Name, req.Phone)
	if err != nil {
		if mgo.IsDup(err) {
			return c.JSON(http.StatusOK, Resp(common.ErrExist))
		} else {
			c.Logger().Error("[New account]", err)
			return c.JSON(http.StatusOK, Resp(common.ErrMongoDB))
		}
	}

	return c.JSON(http.StatusOK, Resp(common.RespSuccess, nil))
}

// 修改状态
func (*accountHandler) ModifyState(c echo.Context) error {
	var req struct {
		Phone string
	}
	req.Phone = c.FormValue("phone")

	if !util.PhoneNum(req.Phone) {
		c.Logger().Error("[Validate]")
		return c.JSON(http.StatusOK, Resp(common.ErrValidate))
	}

	err := model.AccountService.ModifyState(req.Phone)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusOK, Resp(common.ErrNotFound))
		}
		c.Logger().Error("[ModifyState]", err)
		return c.JSON(http.StatusOK, Resp(common.ErrMongoDB))
	}

	return c.JSON(http.StatusOK, Resp(common.RespSuccess, nil))
}

// 改变进出场状态
//func (*accountHandler) InOut(c echo.Context) error {
//	var req struct {
//		Phone string `validate:"numeric,len=11"`
//	}
//
//	req.Phone = c.FormValue("phone")
//
//	if err := c.Validate(&req); err != nil {
//		c.Logger().Error("[Validate]", err)
//		return c.JSON(http.StatusOK, Resp()(common.RespFailed, common.ErrValidate))
//	}
//
//	err := model.AccountService.InOut(req.Phone)
//	if err != nil {
//		c.Logger().Error("[ModifyState]", err)
//		return c.JSON(http.StatusOK, Resp()(common.RespFailed, common.ErrMongo))
//	}
//
//	return c.JSON(http.StatusOK, Resp(common.RespSuccess))
//}

// 信息查询
func (*accountHandler) Info(c echo.Context) error {
	var req struct {
		Phone string `validate:"numeric,len=11"`
	}

	req.Phone = c.FormValue("phone")

	if err := c.Validate(&req); err != nil {
		c.Logger().Error("[Validate]", err)
		return c.JSON(http.StatusOK, Resp(common.ErrValidate))
	}

	a, err := model.AccountService.Info(req.Phone)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusOK, Resp(common.ErrNotFound))
		}
		c.Logger().Error("[Info]", err)
		return c.JSON(http.StatusOK, Resp(common.ErrMongoDB))
	}

	return c.JSON(http.StatusOK, Resp(common.RespSuccess, a))
}

// 用户列表
func (*accountHandler) List(c echo.Context) error {
	a, err := model.AccountService.All()
	if err != nil {
		c.Logger().Error("[List]", err)
		return c.JSON(http.StatusOK, Resp(common.ErrMongoDB))
	}

	return c.JSON(http.StatusOK, Resp(common.RespSuccess, a))
}

// 充值与支付
func (*accountHandler) Recharge(c echo.Context) error {
	var req struct {
		Phone string `validate:"numeric,len=11"`
		Sum   int
	}

	req.Phone = c.FormValue("phone")
	req.Sum, _ = strconv.Atoi(c.FormValue("sum"))

	if req.Sum < 1 {
		return c.JSON(http.StatusOK, Resp(common.ErrParam))
	}

	if err := c.Validate(&req); err != nil {
		c.Logger().Error("[Validate]", err)
		return c.JSON(http.StatusOK, Resp(common.ErrValidate))
	}

	balance, err := model.AccountService.Deal(req.Phone, req.Sum)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusOK, Resp(common.ErrNotFound))
		}
		if err.Error() == common.RespText(common.ErrBalance) {
			c.Logger().Error("[Recharge]", err)
			return c.JSON(http.StatusOK, Resp(common.ErrBalance))
		} else {
			c.Logger().Error("[Recharge]", err)
			return c.JSON(http.StatusOK, Resp(common.ErrDeal))
		}
	}

	return c.JSON(http.StatusOK, Resp(common.RespSuccess, map[string]int{"balance": balance}))
}
