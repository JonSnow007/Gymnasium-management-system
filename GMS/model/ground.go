/*
 * Revision History:
 *     Initial: 2018/05/22        Chen Yanchen
 */

package model

type Ground struct {
	Id int `bson:"_id,omitempty"`
	State int`bson:"State"`// 状态：0.不可用

	size
}

type size struct {
	Wide   int `bson:"Wide"`
	Length int `bson:"Length"`
}
