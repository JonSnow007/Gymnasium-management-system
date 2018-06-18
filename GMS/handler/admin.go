/*
 * Revision History:
 *     Initial: 2018/05/21        Chen Yanchen
 */

package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"

	"github.com/JonSnow007/Gymnasium-management-system/GMS/common"
	"github.com/JonSnow007/Gymnasium-management-system/GMS/model"
	"github.com/JonSnow007/Gymnasium-management-system/GMS/util"
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
		c.Logger().Error("[Parameter]", err)
		return c.JSON(http.StatusOK, Resp(common.ErrParam))
	}

	if err = c.Validate(&req); err != nil {
		c.Logger().Error("[Validate]", err)
		return c.JSON(http.StatusOK, Resp(common.ErrValidate))
	}

	id, err := model.AdminService.New(req.Name, req.Pwd)
	if err != nil {
		c.Logger().Error("[New admin]", err)
		return c.JSON(http.StatusOK, Resp(common.ErrMongoDB))
	}

	return c.JSON(http.StatusOK, Resp(common.RespSuccess, map[string]string{common.RespKeyId: id}))
}

func (*adminHandler) Login(c echo.Context) error {
	var req struct {
		Name string `json:"name" validate:"min=3,max=16"`
		Pwd  string `json:"pwd" validate:"min=6,max=16"`
	}

	if err := c.Bind(&req); err != nil {
		c.Logger().Error("[Parameter]", err)
		return c.JSON(http.StatusOK, Resp(common.ErrParam))
	}

	if err := c.Validate(&req); err != nil {
		c.Logger().Error("[Validate]", err)
		return c.JSON(http.StatusOK, Resp(common.ErrValidate))
	}

	ok, err := model.AdminService.Login(req.Name, req.Pwd)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusOK, Resp(common.ErrNotFound))
		}
		return c.JSON(http.StatusOK, Resp(common.ErrMongoDB))
	}
	func(c echo.Context) {
		cookie := new(http.Cookie)
		cookie.Name = "admin"
		cookie.Value = req.Name
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)
	}(c)

	// todo: get session
	util.GetSession(c, req.Name)
	if ok == true {
		return c.JSON(http.StatusOK, Resp(common.RespSuccess, nil))
	} else {
		return c.JSON(http.StatusOK, Resp(common.ErrAccount))
	}
}

func (*adminHandler) Logout(c echo.Context) (err error) {
	// todo: clear session
	util.ClearSession(c)
	return
}
