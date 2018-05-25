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

	"github.com/JonSnow47/Gymnasium-management-system/GMS/common"
	"github.com/JonSnow47/Gymnasium-management-system/GMS/db"
	"github.com/JonSnow47/Gymnasium-management-system/GMS/util"
)

const collectionAccount = "account"

type Account struct {
	Phone   string    `bson:"_id"`
	Name    string    `bson:"name"`
	Balance float64   `bson:"balance"` // 余额
	Active  bool      `bson:"active"`  // 用户是否在场
	State   bool      `bson:"state"`   // 身份：0.不可用 1.正常使用
	Created time.Time `bson:"created"`
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

	err := con.C.Find(bson.M{"phone": phone}).One(&a)
	if err != nil {
		return err
	}

	err = con.C.Update(bson.M{"active": a.Active}, bson.M{"$set": bson.M{"active": !a.Active}})
	return err
}

// 修改状态
func (*accountServiceProvide) ModifyState(phone string) error {
	con := conAccount()
	defer con.S.Close()

	var a Account

	err := con.C.Find(bson.M{"_id": phone}).Select(bson.M{"state": 1}).One(&a)
	if err != nil {
		return err
	}
	err = con.C.Update(bson.M{"_id": phone}, bson.M{"$set": bson.M{"state": !a.State}})
	return err
}

// 查询信息
func (*accountServiceProvide) Info(phone string) (a *Account, err error) {
	con := conAccount()
	defer con.S.Close()

	err = con.C.Find(bson.M{"_id": phone}).One(&a)
	if err != nil {
		return nil, err
	}

	return
}

// All
func (*accountServiceProvide) All() (a []Account, err error) {
	con := conAccount()
	defer con.S.Close()

	err = con.C.Find(nil).All(&a)
	return
}

// Deal recharge or pay.
func (*accountServiceProvide) Deal(phone string, money float64) (float64, error) {
	con := conAccount()
	defer con.S.Close()

	var a Account
	err := con.C.Find(bson.M{"_id": phone}).Select(bson.M{"balance": 1}).One(&a)
	if err != nil {
		return 0, err
	}

	if a.Balance+money < 0 {
		return 0, errors.New(common.ErrBalance)
	}

	err = con.C.Update(bson.M{"_id": phone, "state": true}, bson.M{"$inc": bson.M{"balance": money}})
	return a.Balance + money, err
}
