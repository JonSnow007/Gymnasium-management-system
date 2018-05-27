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

const collectionGym = "gym"

type Gym struct {
	Id    int    `bson:"_id"`
	Name  string `bson:"Name"`
	IsUse bool   `bson:"IsUse"` // 是否正在使用
	State bool   `bson:"State"` // 状态：0.不可用
}

type gymServiceProvide struct{}

var GymService *gymServiceProvide

func conGym() db.Connection {
	con := db.Connect(collectionGym)
	con.C.EnsureIndex(mgo.Index{
		Key:        []string{"_id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	})
	return con
}

// 新建场馆
func (*gymServiceProvide) New(name string) (err error) {
	con := conGym()
	defer con.S.Close()

	var gym = &Gym{
		Name:  name,
		State: true,
	}

	gym.Id, err = con.C.Find(nil).Count()
	if err != nil {
		return err
	}
	err = con.C.Insert(gym)

	return err
}

// 场馆信息
func (*gymServiceProvide) Info(id int) (g *Gym, err error) {
	con := conGym()
	defer con.S.Close()

	err = con.C.Find(bson.M{"_id": id}).One(&g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

// 场馆列表
func (*gymServiceProvide) List() (g []Gym, err error) {
	con := conGym()
	defer con.S.Close()

	err = con.C.Find(nil).All(&g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

// 修改场馆状态
func (*gymServiceProvide) State(id int) (err error) {
	con := conGym()
	defer con.S.Close()

	var g Gym
	err = con.C.Find(bson.M{"_id": id}).Select(bson.M{"State": 1}).One(&g)
	if err != nil {
		return err
	}
	err = con.C.Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"State": !g.State}})

	return err
}

// 场馆是否正在使用
func (*gymServiceProvide) IsUse(id int) (err error) {
	con := conGym()
	defer con.S.Close()

	var g Gym
	err = con.C.Find(bson.M{"_id": id}).Select(bson.M{"IsUse": 1}).One(&g)
	if err != nil {
		return err
	}
	err = con.C.Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"IsUse": !g.IsUse}})

	return err
}
