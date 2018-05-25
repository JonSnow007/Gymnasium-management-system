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
	RespStatus  = true
	RespSuccess = "success"
	RespFailed  = "failed"
)

// response error
const (
	ErrParam    = "Parameter" // 参数错误
	ErrValidate = "Validate"  // 参数验证失败

	ErrExist         = "Exist"                      // 已存在
	ErrNotFound      = "Not found"                  // 未查询到数据
	ErrLogin         = "Incorrect name or password" // 密码错误
	ErrLoginRequired = "Login required"             // 未登录
	ErrPerm          = "Invalid permission"         // 权限错误

	ErrSession = "Session error" // Session 错误

	ErrDeal    = "Deal failed"     // 交易失败
	ErrBalance = "Lack of balance" // 余额不足

	ErrMongo = "Mongodb error" // MongoDB 错误
)
