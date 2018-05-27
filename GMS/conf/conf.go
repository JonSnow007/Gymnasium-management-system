/*
 * Revision History:
 *     Initial: 2018/05/21        Chen Yanchen
 */

package conf

import (
	"log"

	"github.com/JonSnow47/Gymnasium-management-system/GMS/util"
)

const (
	MongoURL = "localhost:27017"
	MongoDB  = "GMS"
)

type Gms struct {
	Price float32
}

type Mgo struct {
	Database string
	URL      string
}

type Config struct {
	Mod string
	Mgo *Mgo
	Gms *Gms
}

var Conf Config

func init() {
	err := util.ParseConf("conf/conf.json", &Conf)
	if err != nil {
		log.Println("[Parse configuraion]", err)
	}
	//fmt.Println(Conf)
}
