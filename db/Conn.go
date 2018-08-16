package db

import (
	"database/sql"
	"fmt"
	"position_postgre/conf"
	"position_postgre/tools"

	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	var (
		host     = conf.ConfigContext.DBUrl
		port     = conf.ConfigContext.DBPort
		user     = conf.ConfigContext.DBUser
		password = conf.ConfigContext.DBPwd
		dbname   = conf.ConfigContext.DBName
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	tools.PanicError(err)

	err = DB.Ping()
	tools.PanicError(err)

	if err == nil{
		log.Println("connect db success")
	}
}
