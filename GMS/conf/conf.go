/*
 * Revision History:
 *     Initial: 2018/05/21        Chen Yanchen
 */

package conf

const (
	MongoURL = "localhost:27017"
	MongoDB  = "GMS"

	Price = 20
)

type Gms struct {
	Price int
}

type Mgo struct {
	Database string
	URL      string
}

type Config struct {
	Mod string
	Mgo Mgo
	Gms Gms
}

var Conf = &Config{
	Mod: "dev",
	Mgo: Mgo{
		Database: MongoDB,
		URL:      MongoURL,
	},
	Gms: Gms{
		Price: Price,
	},
}

//var Conf Config

//func init() {
//	err := util.ParseConf("./conf/conf.json", &Conf)
//	if err != nil {
//		log.Println("[Parse configuraion]", err)
//	}
//}
