/*
 * Revision History:
 *     Initial: 2018/05/22        Chen Yanchen
 */

package util

import (
	"regexp"
)

// 验证手机号码是否合格
func PhoneNum(phone string) bool {
	if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, phone); !m {
		return false
	}
	return true
}
