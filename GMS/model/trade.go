/*
 * Revision History:
 *     Initial: 2018/05/22        Chen Yanchen
 */

package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/JonSnow47/Gymnasium-management-system/GMS/db"
)

const collectionTrade = "trade"

type Bill struct {
	Id bson.ObjectId `_id,omitempty`
}

type billServiceProvide struct{}

var Service *accountServiceProvide

func conTrade() db.Connection {
	con := db.Connect(collectionTrade)
	con.C.EnsureIndex(mgo.Index{
		Key:        []string{"_id", "Phone"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	})
	return con
}
