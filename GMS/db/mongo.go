/*
 * Revision History:
 *     Initial: 2018/05/21        Chen Yanchen
 */

package db

import (
	"gopkg.in/mgo.v2"

	"github.com/JonSnow007/Gymnasium-management-system/GMS/conf"
)

type Connection struct {
	S *mgo.Session
	D *mgo.Database
	C *mgo.Collection
}

// 连接 MongoDB
func Connect(collection string) (con Connection) {
	var err error
	url := conf.Conf.Mgo.URL + "/" + conf.Conf.Mgo.Database

	con.S, err = mgo.Dial(url)
	if err != nil {
		panic(err)
	}

	con.S.SetMode(mgo.Monotonic, true)

	con.D = con.S.DB(conf.Conf.Mgo.Database)
	con.C = con.D.C(collection)
	return
}
