/*
 * Revision History:
 *     Initial: 2018/05/21        Chen Yanchen
 */

package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/JonSnow47/Gymnasium-management-system/GMS/common"
	"github.com/JonSnow47/Gymnasium-management-system/GMS/model"
)

type adminHandler struct{}

var Admin *adminHandler

func (*adminHandler) New(c echo.Context) error {
	var (
		req struct {
			Name string `json:"name" validate:"min=3,max=16"`
			Pwd  string `json:"pwd" validate:"min=6,max=16"`
		}
		err error
	)

	if err = c.Bind(&req); err != nil {
		c.Logger().Info("[Parameter]", err)
		return c.JSON(http.StatusOK, RespData(common.RespFailed, common.ErrParam))
	}

	if err = c.Validate(&req); err != nil {
		c.Logger().Info("[Validate]", err)
		return c.JSON(http.StatusOK, RespData(common.RespFailed, common.ErrValidate))
	}

	id, err := model.AdminService.New(req.Name, req.Pwd)
	if err != nil {
		c.Logger().Error("[New admin]", err)
		return c.JSON(http.StatusOK, RespData(common.RespFailed, common.ErrMongo))
	}

	return c.JSON(http.StatusOK, RespId(common.RespSuccess, id))
}

func (*adminHandler) Login(c echo.Context) error {
	var req struct {
		Name string `json:"name" validate:"min=3,max=16"`
		Pwd  string `json:"pwd" validate:"min=6,max=16"`
	}

	if err := c.Bind(&req); err != nil {
		c.Logger().Info("[Parameter]", err)
		return c.JSON(http.StatusOK, RespData(common.RespFailed, common.ErrParam))
	}

	if err := c.Validate(&req); err != nil {
		c.Logger().Info("[Validate]", err)
		return c.JSON(http.StatusOK, RespData(common.RespFailed, common.ErrValidate))
	}

	ok, err := model.AdminService.Login(req.Name, req.Pwd)
	if err == nil {
		if ok == true {
			return c.JSON(http.StatusOK, RespData(common.RespSuccess))
		} else {
			return c.JSON(http.StatusOK, RespData(common.RespSuccess, common.ErrLogin))
		}
	}

	c.Logger().Error("[Admin Login]", err)
	return c.JSON(http.StatusOK, RespData(common.RespFailed, common.ErrMongo))
}