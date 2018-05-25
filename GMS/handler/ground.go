/*
 * Revision History:
 *     Initial: 2018/05/24        Chen Yanchen
 */

package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"

	"github.com/JonSnow47/Gymnasium-management-system/GMS/common"
	"github.com/JonSnow47/Gymnasium-management-system/GMS/model"
)

type groundHandler struct{}

var Ground *groundHandler

// 新建场地
func (*groundHandler) New(c echo.Context) error {
	var (
		err error
		req struct {
			Name string `json:"name" validate:"required,min=1,max=20"`
		}
	)

	req.Name = c.FormValue("name")

	if err = c.Validate(&req); err != nil {
		c.Logger().Error("[Validate]", err)
		return c.JSON(http.StatusOK, RespData(common.RespFailed, common.ErrValidate))
	}

	err = model.GroundService.New(req.Name)
	if err != nil {
		c.Logger().Error("[New account]", err)
		return c.JSON(http.StatusOK, RespData(common.RespFailed, common.ErrMongo))
	}

	return c.JSON(http.StatusOK, Resp(common.RespSuccess))
}

// 场地信息
func (*groundHandler) Info(c echo.Context) error {
	var (
		err error
		req struct {
			Id int
		}
	)

	req.Id, err = strconv.Atoi(c.FormValue("id"))
	if err != nil {
		c.Logger().Error("[Validate]", err)
		return c.JSON(http.StatusOK, RespData(common.RespFailed, common.ErrValidate))
	}

	g, err := model.GroundService.Info(req.Id)
	if err != nil {
		if err == mgo.ErrNotFound {
			c.Logger().Error("[Info]", err)
			return c.JSON(http.StatusOK, RespData(common.RespFailed, common.ErrNotFound))
		}
		c.Logger().Error("[Info]", err)
		return c.JSON(http.StatusOK, RespData(common.RespFailed, common.ErrMongo))
	}

	return c.JSON(http.StatusOK, RespData(common.RespSuccess, g))
}

// 场地列表
func (*groundHandler) List(c echo.Context) error {

	g, err := model.GroundService.List()
	if err != nil {
		c.Logger().Error("[List]", err)
		return c.JSON(http.StatusOK, RespData(common.RespFailed, common.ErrMongo))
	}

	return c.JSON(http.StatusOK, RespData(common.RespSuccess, g))
}

// 修改场地状态
func (*groundHandler) ModifyState(c echo.Context) error {
	var (
		err error
		req struct {
			Id int
		}
	)

	req.Id, err = strconv.Atoi(c.FormValue("id"))
	if err != nil {
		c.Logger().Error("[Validate]", err)
		return c.JSON(http.StatusOK, RespData(common.RespFailed, common.ErrValidate))
	}

	err = model.GroundService.State(req.Id)
	if err != nil {
		c.Logger().Error("[ModifyState]", err)
		return c.JSON(http.StatusOK, RespData(common.RespFailed, common.ErrMongo))
	}

	return c.JSON(http.StatusOK, Resp(common.RespSuccess))
}