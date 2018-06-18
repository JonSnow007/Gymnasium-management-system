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
	"fmt"
)

const collectionAccount = "account"

// Account represents the account information.
type Account struct {
	Phone    string    `bson:"_id"`
	Name     string    `bson:"Name"`
	Balance  int       `bson:"Balance"`  // 余额
	Recorded int       `bson:"Recorded"` // 充值总金额
	Active   bool      `bson:"Active"`   // 用户是否在场
	State    bool      `bson:"State"`    // 身份：0.不可用 1.正常使用
	Created  time.Time `bson:"Created"`
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

// New account.
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
	fmt.Println(err)
	if err != nil {
		return err
	}

	return nil
}

// InOut represents if account in gym.
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

// ModifyState modify account state(usable/unusable).
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

// Info represents get a account info.
func (*accountServiceProvide) Info(phone string) (a *Account, err error) {
	con := conAccount()
	defer con.S.Close()

	err = con.C.Find(bson.M{"_id": phone}).One(&a)
	if err != nil {
		return
	}

	return
}

// All account list.
func (*accountServiceProvide) All() (a []*Account, err error) {
	con := conAccount()
	defer con.S.Close()

	err = con.C.Find(nil).Sort("-Created").All(&a)
	return
}

// Deal recharge or pay.
func (*accountServiceProvide) Deal(phone string, money int) (int, error) {
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

	if money > 0 {
		err = con.C.Update(bson.M{"_id": phone, "State": true}, bson.M{"$inc": bson.M{
			"Balance":  money,
			"Recorded": money,
		}})
		if err != nil {
			return 0, err
		}
	} else {
		err = con.C.Update(bson.M{"_id": phone, "State": true}, bson.M{"$inc": bson.M{"Balance": money}})
	}
	return a.Balance + money, err
}

// Recorded represents all recorded.
func (*accountServiceProvide) Recorded() (recorded int, err error) {
	con := conAccount()
	defer con.S.Close()

	var accounts []*Account
	err = con.C.Find(nil).Select("Recorded").All(&accounts)
	if err != nil {
		return
	}
	for i, _ := range accounts {
		recorded += accounts[i].Recorded
	}
	return
}
