/*
 * Revision History:
 *     Initial: 2018/05/21        Chen Yanchen
 */

package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/JonSnow47/Gymnasium-management-system/GMS/common"
)

func Resp(status string) map[string]string {
	return map[string]string{common.RespKeyStatus: status}
}

// 对请求结果的状态和结果排版
func RespData(status string, data interface{}) map[string]interface{} {
	if data == nil {
		return map[string]interface{}{common.RespKeyStatus: status}
	}
	if status == common.RespFailed {
		return map[string]interface{}{common.RespKeyStatus: status, common.RespKeyErr: data}
	}
	return map[string]interface{}{common.RespKeyStatus: status, common.RespKeyData: data}
}

func RespId(status string, id interface{}) map[string]interface{} {
	return map[string]interface{}{common.RespKeyStatus: status, common.RespKeyId: id}
}

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Gymnasium-management-system!")
}
