/*
 * Revision History:
 *     Initial: 2018/05/22        Chen Yanchen
 */

package model

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/JonSnow007/Gymnasium-management-system/GMS/db"
	"github.com/JonSnow007/Gymnasium-management-system/GMS/util"
)

const collectionAccount = "account"

type Account struct {
	Phone   string    `bson:"_id"`
	Name    string    `bson:"Name"`
	Balance float32   `bson:"Balance"` // 余额
	Active  bool      `bson:"Active"`  // 用户是否在场
	State   bool      `bson:"State"`   // 身份：0.不可用 1.正常使用
	Created time.Time `bson:"Created"`
}

type accountServiceProvide struct{}

var AccountService *accountServiceProvide

func conAccount() db.Connection {
	con := db.Connect(collectionAccount)
	con.C.EnsureIndex(mgo.Index{
		Key:        []string{"_id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	})
	return con
}

func (*accountServiceProvide) New(name, phone string) error {
	con := conAccount()
	defer con.S.Close()

	if !util.PhoneNum(phone) {
		return errors.New("Incorrect phone num")
	}

	var account = &Account{
		Name:    name,
		Phone:   phone,
		Balance: 0,
		Created: time.Now(),
		State:   true,
	}

	err := con.C.Insert(account)
	if err != nil {
		return err
	}

	return nil
}

// ModifyPhone forbidden
func (*accountServiceProvide) ModifyPhone(old, new string) error {
	con := conAccount()
	defer con.S.Close()

	if !util.PhoneNum(new) {
		return errors.New("Incorrect phone num")
	}

	err := con.C.Update(bson.M{"_id": old}, bson.M{"$set": bson.M{"_id": new}})
	return err
}

// 修改是否在场状态
func (*accountServiceProvide) InOut(phone string) error {
	con := conAccount()
	defer con.S.Close()

	var a Account

	err := con.C.Find(bson.M{"_id": phone}).One(&a)
	if err != nil {
		return err
	}

	err = con.C.Update(bson.M{"_id": phone}, bson.M{"$set": bson.M{"Active": !a.Active}})
	return err
}

// 修改状态
func (*accountServiceProvide) ModifyState(phone string) error {
	con := conAccount()
	defer con.S.Close()

	var a Account

	err := con.C.Find(bson.M{"_id": phone}).Select(bson.M{"State": 1}).One(&a)
	if err != nil {
		return err
	}
	err = con.C.Update(bson.M{"_id": phone}, bson.M{"$set": bson.M{"State": !a.State}})
	return err
}

// 查询信息
func (*accountServiceProvide) Info(phone string) (a *Account, err error) {
	con := conAccount()
	defer con.S.Close()

	err = con.C.Find(bson.M{"_id": phone}).One(&a)
	if err != nil {
		return
	}

	return
}

// All
func (*accountServiceProvide) All() (a []*Account, err error) {
	con := conAccount()
	defer con.S.Close()

	err = con.C.Find(nil).Sort("-Created").All(&a)
	return
}

// Deal recharge or pay.
func (*accountServiceProvide) Deal(phone string, money float32) (float32, error) {
	con := conAccount()
	defer con.S.Close()

	var a Account
	err := con.C.Find(bson.M{"_id": phone}).Select(bson.M{"Balance": 1}).One(&a)
	if err != nil {
		return 0, err
	}

	if a.Balance+money < 0 {
		return 0, errors.New("Lack of balance")
	}

	err = con.C.Update(bson.M{"_id": phone, "State": true}, bson.M{"$inc": bson.M{"Balance": money}})
	return a.Balance + money, err
}
