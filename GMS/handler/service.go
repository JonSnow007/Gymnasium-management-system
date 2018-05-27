/*
 * Revision History:
 *     Initial: 2018/05/27        Chen Yanchen
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

type serviceProdive struct{}

var Service *serviceProdive

// Home service.
func (*serviceProdive) Home(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Gymnasium-management-system!")
}

// In account come in Gymnasium service.
func (*serviceProdive) In(c echo.Context) error {
	var (
		req struct {
			Phone string
			Gid   int
		}
		err error
	)
	req.Phone = c.FormValue("phone")
	req.Gid, err = strconv.Atoi(c.FormValue("gid"))
	if err != nil || !util.PhoneNum(req.Phone) {
		c.Logger().Error("[Validate]", err)
		return c.JSON(http.StatusOK, Resp(common.RespFailed, common.ErrValidate))
	}
	// ensure account info
	account, err := model.AccountService.Info(req.Phone)
	if err != nil {
		c.Logger().Error("[In]", err)
		return c.JSON(http.StatusOK, Resp(common.RespFailed, common.ErrMongo))
	}
	if account.Balance <= 0 {
		c.Logger().Error("[In]", common.ErrBalance)
		return c.JSON(http.StatusOK, Resp(common.RespFailed, common.ErrBalance))
	}
	if account.Active == true {
		c.Logger().Error("[In]", common.ErrGymUsing)
		return c.JSON(http.StatusOK, Resp(common.RespFailed, common.ErrGymUsing))
	}
	// ensure gym info
	gym, err := model.GymService.Info(req.Gid)
	if err != nil {
		c.Logger().Error("[In]", err)
		return c.JSON(http.StatusOK, Resp(common.RespFailed, common.ErrMongo))
	}
	if gym.State == false {
		c.Logger().Error("[In]", common.ErrForbidden)
		return c.JSON(http.StatusOK, Resp(common.RespFailed, common.ErrForbidden))
	}
	if gym.IsUse == true {
		c.Logger().Error("[In]", common.ErrGymUsing)
		return c.JSON(http.StatusOK, Resp(common.RespFailed, common.ErrGymUsing))
	}

	err = model.BillService.New(req.Phone, req.Gid)
	if err != nil {
		c.Logger().Error("[In]", common.ErrGymUsing)
		return c.JSON(http.StatusOK, Resp(common.RespFailed, common.ErrMongo))
	}

	err = model.AccountService.InOut(req.Phone)
	err = model.GymService.IsUse(req.Gid)
	return c.JSON(http.StatusOK, Resp(common.RespSuccess, nil))
}

// Out account out of gymnasium service.
func (*serviceProdive) Out(c echo.Context) (err error) {
	phone := c.FormValue("phone")

	if !util.PhoneNum(phone) {
		c.Logger().Error("[Validate]", err)
		return c.JSON(http.StatusOK, Resp(common.RespFailed, common.ErrValidate))
	}

	b, err := model.BillService.Clearing(phone)
	if err != nil {
		c.Logger().Error("[Out]", err)
		return c.JSON(http.StatusOK, Resp(common.RespFailed, common.ErrDeal))
	}

	// 余额
	_, err = model.AccountService.Deal(phone, -b.Consume)
	err = model.AccountService.InOut(b.Phone)
	err = model.GymService.IsUse(b.Gid)

	return c.JSON(http.StatusOK, Resp(common.RespSuccess, &b))
}
