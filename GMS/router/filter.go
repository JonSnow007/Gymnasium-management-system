/*
 * Revision History:
 *     Initial: 2018/06/12        Chen Yanchen
 */

package router

import (
	"fmt"

	"github.com/labstack/echo"
)

var RouterFilter map[string]interface{}

func init() {
	RouterFilter = make(map[string]interface{})
	RouterFilter["/api/home"] = struct{}{}

	RouterFilter["/api/admin/new"] = struct{}{}
	RouterFilter["/api/admin/login"] = struct{}{}
}

func LoginFilter(ctx echo.Context) {
	if _, ok := RouterFilter[ctx.Request().RequestURI]; !ok {
		for _, cookie := range ctx.Cookies() {
			fmt.Println(cookie.Name)
			fmt.Println(cookie.Value)
		}
	}
}
