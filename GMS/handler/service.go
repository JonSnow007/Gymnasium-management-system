/*
 * Revision History:
 *     Initial: 2018/05/27        Chen Yanchen
 */

package handler

import (
	"github.com/labstack/echo"
	"net/http"

	"github.com/JonSnow007/Gymnasium-management-system/GMS/common"
	"github.com/JonSnow007/Gymnasium-management-system/GMS/model"
	"github.com/JonSnow007/Gymnasium-management-system/GMS/util"
)

type serviceProdive struct{}

var Service *serviceProdive

// Home service.
func (*serviceProdive) Home(c echo.Context) error {
	return c.String(http.StatusOK, "Thanks for use Gymnasium-management-system!\n体育场馆管理系统已启动!")
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

	//req.Phone = c.FormValue("phone")
	//req.Gid, err = strconv.Atoi(c.FormValue("gid"))
	if err = c.Bind(&req); err != nil {
		c.Logger().Error("[Bind]", err)
		return c.JSON(http.StatusOK, Resp(common.ErrParam))
	}

	if err != nil || !util.PhoneNum(req.Phone) {
		c.Logger().Error("[Validate]", err)
		return c.JSON(http.StatusOK, Resp(common.ErrValidate))
	}
	// ensure account info
	account, err := model.AccountService.Info(req.Phone)
	if err != nil {
		c.Logger().Error("[In]", err)
		return c.JSON(http.StatusOK, Resp(common.ErrMongoDB))
	}
	if account.Balance <= 0 {
		c.Logger().Error("[In]", common.ErrBalance)
		return c.JSON(http.StatusOK, Resp(common.ErrBalance))
	}
	if account.Active == true {
		c.Logger().Error("[In]", common.ErrGymUsing)
		return c.JSON(http.StatusOK, Resp(common.ErrGymUsing))
	}
	// ensure gym info
	gym, err := model.GymService.Info(req.Gid)
	if err != nil {
		c.Logger().Error("[In]", err)
		return c.JSON(http.StatusOK, Resp(common.ErrMongoDB))
	}
	if gym.State == false {
		c.Logger().Error("[In]", common.ErrForbidden)
		return c.JSON(http.StatusOK, Resp(common.ErrForbidden))
	}
	if gym.IsUse == true {
		c.Logger().Error("[In]", common.ErrGymUsing)
		return c.JSON(http.StatusOK, Resp(common.ErrGymUsing))
	}

	err = model.BillService.New(req.Phone, req.Gid)
	if err != nil {
		c.Logger().Error("[In]", common.ErrGymUsing)
		return c.JSON(http.StatusOK, Resp(common.ErrMongoDB))
	}

	err = model.AccountService.InOut(req.Phone)
	err = model.GymService.IsUse(req.Gid)
	return c.JSON(http.StatusOK, Resp(common.RespSuccess))
}

// Out account out of gymnasium service.
func (*serviceProdive) Out(c echo.Context) (err error) {
	var req struct {
		Phone string `json:"phone" validate:"required,len=11"`
	}
	//req.Phone = c.FormValue("phone")
	if err = c.Bind(&req); err != nil {
		c.Logger().Error("[Bind]", err)
		return c.JSON(http.StatusOK, Resp(common.ErrParam))
	}

	if !util.PhoneNum(req.Phone) {
		c.Logger().Error("[Validate]", err)
		return c.JSON(http.StatusOK, Resp(common.ErrValidate))
	}

	b, err := model.BillService.Clearing(req.Phone)
	if err != nil {
		c.Logger().Error("[Out]", err)
		return c.JSON(http.StatusOK, Resp(common.ErrDeal))
	}

	// 余额
	balance, err := model.AccountService.Deal(req.Phone, -b.Consume)
	err = model.AccountService.InOut(b.Phone)
	err = model.GymService.IsUse(b.Gid)

	return c.JSON(http.StatusOK, Resp(common.RespSuccess, &b, map[string]int{"balance": balance}))
}
