/*
 * Revision History:
 *     Initial: 2018/05/21        Chen Yanchen
 */

package util

import (
	"encoding/json"
	"io/ioutil"
)

// 解析 .json 类型配置文件
func ParseConf(filename string, v interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, v)

	return err
}
