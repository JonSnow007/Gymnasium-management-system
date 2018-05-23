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

	"github.com/JonSnow47/Gymnasium-management-system/GMS/db"
	"github.com/JonSnow47/Gymnasium-management-system/GMS/util"
	"github.com/JonSnow47/Gymnasium-management-system/GMS/common"
)

const collectionAccount = "account"

type Account struct {
	Id      bson.ObjectId `bson:"_id,omitempty"`
	Name    string        `bson:"Name"`
	Phone   string        `bson:"Phone"`
	Balance float64       `bson:"Balance"` // 余额
	Active  bool          `bson:"Active"`  // 用户是否在场
	State   bool          `bson:"Status"`  // 身份：0.不可用 1.正常使用
	Created time.Time     `bson:"Created"`
}

type accountServiceProvide struct{}

var AccountService *accountServiceProvide

func conAccount() db.Connection {
	con := db.Connect(collectionAccount)
	con.C.EnsureIndex(mgo.Index{
		Key:        []string{"_id", "Phone"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	})
	return con
}

func (*accountServiceProvide) New(name, phone string) (bson.ObjectId, error) {
	con := conAccount()
	defer con.S.Close()

	if !util.PhoneNum(phone) {
		return "", errors.New("Incorrect phone num")
	}

	var account = &Account{
		//Id:      bson.NewObjectId(),
		Name:    name,
		Phone:   phone,
		Balance: 0,
		Created: time.Now(),
		State:   true,
	}

	err := con.C.Insert(account)
	if err != nil {
		return "", err
	}

	return account.Id, nil
}

func (*accountServiceProvide) ModifyPhone(old, new string) error {
	con := conAccount()
	defer con.S.Close()

	if !util.PhoneNum(new) {
		return errors.New("Incorrect phone num")
	}

	err := con.C.Update(bson.M{"Phone": old, "State": true}, bson.M{"$set": bson.M{"Phone": new}})
	return err
}

// 修改是否在场状态
func (*accountServiceProvide) InOut(phone string) error {
	con := conAccount()
	defer con.S.Close()

	var a Account

	err := con.C.Find(bson.M{"Phone": phone}).One(&a)
	if err != nil {
		return err
	}

	err = con.C.Update(bson.M{"In": a.Active}, bson.M{"$set": bson.M{"In": !a.Active}})
	return err
}

// 修改状态
func (*accountServiceProvide) ModifyState(phone string) error {
	con := conAccount()
	defer con.S.Close()

	var a Account

	err := con.C.Find(bson.M{"Phone": phone}).One(&a)
	if err != nil {
		return err
	}

	err = con.C.Update(bson.M{"State": a.Active}, bson.M{"$set": bson.M{"State": !a.Active}})
	return err
}

// 查询信息
func (*accountServiceProvide) Info(phone string) (a *Account, err error) {
	con := conAccount()
	defer con.S.Close()

	err = con.C.Find(bson.M{"Phone": phone}).One(&a)
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

// Deal
func (*accountServiceProvide) Deal(phone string, money int) error {
	con := conAccount()
	defer con.S.Close()

	var sum int
	err := con.C.Find(bson.M{"Phone": phone}).Select(bson.M{"Balance": 1}).One(&sum)
	if err != nil {
		return err
	}

	if sum-money < 0 {
		return errors.New(common.ErrBalance)
	}

	err = con.C.Update(bson.M{"Phone": phone, "State": true}, bson.M{"$inc": bson.M{"Balance": money}})
	return err
}

func Test(phone string) {
	con := conAccount()
	defer con.S.Close()

	var res string
	con.C.Find(bson.M{"Phone": phone}).Select(bson.M{"Name": 1}).One(&res)

	print(res)
}
