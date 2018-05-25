/*
 * Revision History:
 *     Initial: 2018/05/25        Chen Yanchen
 */

package model

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/JonSnow47/Gymnasium-management-system/GMS/db"
)

const collectionBill = "bill"

type Bill struct {
	Id      bson.ObjectId `bson:"_id"`
	Phone   string        `bson:"phone,len=11"` // 用户 id
	Pid     int           `bson:"pid"`          // 场地 id
	Sum     float64       `bson:"sum"`          // 金额
	State   bool          `bson:"state"`        // true:完成 false:未完成
	Created time.Time     `bson:"created"`      // 创建时间
	Updated time.Time     `bson:"updated"`      // 交易完成时间
}

type billServiceProvide struct{}

var BillService *billServiceProvide

func conBill() db.Connection {
	con := db.Connect(collectionBill)
	con.C.EnsureIndex(mgo.Index{
		Key:        []string{"_id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	})
	return con
}

func (*billServiceProvide) New(phone string, pid int, sum float64) (err error) {
	con := conBill()
	defer con.S.Close()

	var bill = &Bill{
		Id:      bson.NewObjectId(),
		Phone:   phone,
		Pid:     pid,
		Sum:     sum,
		Created: time.Now(),
	}

	err = con.C.Insert(bill)

	return err
}

func (*billServiceProvide) State(id bson.ObjectId) (err error) {
	con := conBill()
	defer con.S.Close()

	err = con.C.Update(bson.M{"_id": id},
		bson.M{"$set": bson.M{"state": true, "updated": time.Now()}})
	if err != nil {
		return err
	}

	return nil
}

func (*billServiceProvide) Info(id string) (b *Bill, err error) {
	con := conBill()
	defer con.S.Close()

	err = con.C.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (*billServiceProvide) ListByPhone(phone string) (b *[]Bill, err error) {
	con := conBill()
	defer con.S.Close()

	err = con.C.Find(bson.M{"phone": phone}).Sort("-created").All(&b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (*billServiceProvide) ListByPid(pid int) (b *[]Bill, err error) {
	con := conBill()
	defer con.S.Close()

	err = con.C.Find(bson.M{"pid": pid}).Sort("-created").All(&b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (*billServiceProvide) List() (b *[]Bill, err error) {
	con := conBill()
	defer con.S.Close()

	err = con.C.Find(nil).Sort("-created").All(&b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
