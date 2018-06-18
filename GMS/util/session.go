/*
 * Revision History:
 *     Initial: 2018/06/12        Chen Yanchen
 */

package util

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

// GetSession to get a new session.
func GetSession(ctx echo.Context, name string) {
	session, err := store.Get(ctx.Request(), "Gymnasium-management-system")
	if err != nil {
		http.Error(ctx.Response().Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["admin"] = name

	session.Save(ctx.Request(), ctx.Response().Writer)
}

// ClearSession clearing all session.
func ClearSession(ctx echo.Context) {
	session, err := store.Get(ctx.Request(), "Gymnasium-management-system")
	if err != nil {
		http.Error(ctx.Response().Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values = make(map[interface{}]interface{})
	sessions.Save(ctx.Request(), ctx.Response().Writer)
}
