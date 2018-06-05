/*
 * Revision History:
 *     Initial: 2018/05/25        Chen Yanchen
 */

package model

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/JonSnow007/Gymnasium-management-system/GMS/conf"
	"github.com/JonSnow007/Gymnasium-management-system/GMS/db"
)

const collectionBill = "bill"

type Bill struct {
	Id       bson.ObjectId `bson:"_id"`
	Phone    string        `bson:"Phone"`    // 用户 id
	Gid      int           `bson:"Gid"`      // 场地 id
	State    bool          `bson:"State"`    // true:完成 false:未完成
	InAt     time.Time     `bson:"InAt"`     // 入场时间
	OutAt    time.Time     `bson:"OutAt"`    // 出场时间
	Duration int           `bson:"Duration"` // 使用时长: min
	Price    int           `bson:"Price"`    // 单价: ￥/h
	Consume  int           `bson:"Consume"`  // 消费: ￥
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

// New bill.
func (*billServiceProvide) New(phone string, gid int) (err error) {
	con := conBill()
	defer con.S.Close()

	var bill = &Bill{
		Id:    bson.NewObjectId(),
		Phone: phone,
		Gid:   gid,
		InAt:  time.Now(),
		Price: conf.Conf.Gms.Price,
	}

	err = con.C.Insert(bill)

	return err
}

// Clearing the bill.
func (*billServiceProvide) Clearing(phone string) (b Bill, err error) {
	con := conBill()
	defer con.S.Close()

	err = con.C.Find(bson.M{"Phone": phone, "State": false}).Sort("-InAt").One(&b)
	if err != nil {
		return
	}

	b.OutAt = time.Now()
	b.Duration = int(time.Since(b.InAt).Minutes())
	if b.Duration%60 > 15 {
		b.Consume = int(b.Duration*b.Price/60 + b.Price)
	} else {
		b.Consume = int(b.Duration * b.Price / 60)
	}
	b.State = true

	err = con.C.Update(bson.M{"_id": b.Id}, &b)
	if err != nil {
		return
	}

	return
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

func (*billServiceProvide) ListByPhone(phone string) (b []*Bill, err error) {
	con := conBill()
	defer con.S.Close()

	err = con.C.Find(bson.M{"Phone": phone}).Sort("-InAt").All(&b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (*billServiceProvide) ListByGid(gid int) (b []*Bill, err error) {
	con := conBill()
	defer con.S.Close()

	err = con.C.Find(bson.M{"Gid": gid}).Sort("-InAt").All(&b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (*billServiceProvide) List() (b []*Bill, err error) {
	con := conBill()
	defer con.S.Close()

	err = con.C.Find(nil).Sort("-InAt").All(&b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (*billServiceProvide) Total() (total int, err error) {
	con := conBill()
	defer con.S.Close()

	var bills []*Bill
	err = con.C.Find(nil).All(&bills)
	if err != nil {
		return
	}

	for i, _ := range bills {
		total += bills[i].Consume
	}

	return
}
