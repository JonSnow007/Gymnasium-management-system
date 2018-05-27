/*
 * Revision History:
 *     Initial: 2018/05/21        Chen Yanchen
 */

package handler

import (
	"github.com/JonSnow47/Gymnasium-management-system/GMS/common"
)

// Resp format the response status and data/error.
func Resp(status string, data interface{}) map[string]interface{} {
	if data == nil {
		return map[string]interface{}{common.RespKeyStatus: status}
	}
	if status == common.RespSuccess {
		return map[string]interface{}{common.RespKeyStatus: status, common.RespKeyData: data}
	}
	return map[string]interface{}{common.RespKeyStatus: status, common.RespKeyErr: data}
}
