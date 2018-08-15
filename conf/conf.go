package conf

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	DBUrl            string
	DBPort           int
	DBUser           string
	DBPwd            string
	ServerPort       string
	DBName           string
	ChangePwdToken   string
	QiNiu_ACCESS_KEY string
	QiNiu_SECRET_KEY string
}

var ConfigContext Config

func ReadConf() {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	ConfigContext = Config{}
	err := decoder.Decode(&ConfigContext)
	if err != nil {
		log.Println("读取配置文件 conf.json 错误错误❌")
	}
}
func init() {
	ReadConf()
	log.Println("read conf ")
}
