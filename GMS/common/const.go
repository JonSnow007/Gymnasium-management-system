/*
 * Revision History:
 *     Initial: 2018/05/22        Chen Yanchen
 */

package common

// response key
const (
	RespKeyStatus = "status"
	RespKeyId     = "id"
	RespKeyToken  = "token"
	RespKeyData   = "data"
	RespKeyErr    = "error"
)

// response status
const (
	RespSuccess = 0
	RespFailed  = 1

	ErrParam    = 411 // ErrInvalidParam - Invalid Parameter
	ErrValidate = 412

	ErrPermission = 421 // ErrPermission - Permission Denied
	ErrForbidden  = 422
	ErrExist      = 423
	ErrNotFound   = 424
	ErrAccount    = 425 // ErrAccount - No This User or Password Error

	ErrInternalServerError = 500 // ErrInternalServerError - Internal error.

	ErrDeal     = 520 // ErrWechatPay - Wechat Pay error.
	ErrBalance  = 521 // ErrWechatAuth - Wechat Auth error.
	ErrGymUsing = 522

	ErrMongoDB = 600 // ErrMongoDB - MongoDB operations error.
	ErrMysql   = 700 // ErrMysql - Mysql operations error.
)

// respText
var respText = map[int]string{
	ErrParam:    "Invalid Parameter", // 参数错误
	ErrValidate: "Validate",          // 参数验证失败

	ErrForbidden:  "Forbidden",                  // 禁止使用
	ErrExist:      "Exist",                      // 已存在
	ErrNotFound:   "Not found",                  // 未查询到数据
	ErrAccount:    "Incorrect name or password", // 密码错误
	ErrPermission: "Invalid permission",         // 权限错误

	ErrInternalServerError: "Internal error",

	ErrMongoDB: "Mongodb error", // MongoDB 错误

	ErrDeal:     "Deal failed",     // 交易失败
	ErrBalance:  "Lack of balance", // 余额不足
	ErrGymUsing: "Using",           // 使用中
}

func RespText(code int) string {
	return respText[code]
}
