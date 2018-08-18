package models

import (
	//"fmt"
	"position_postgre/db"
	//"position_postgre/tools"
	//"log"
	"database/sql"
	"position_postgre/modelStruct"
	"position_postgre/tools"
)

var (
	FIND_USER_SQL = `SELECT id,user_name,nick_name,avatar_url,openid FROM users WHERE openid=$1;`
)

func FindUserByOpenid(openId string) *sql.Rows {

	sqlStatement := FIND_USER_SQL
	rows, err := db.DB.Query(sqlStatement, openId)
	//fmt.Println(err)
	//defer rows.Close()

	tools.PanicError(err)
	return rows
}

func InsertUser(user modelStruct.User) {
	BuildInsertSql("INSERT INTO users", user, db.DB.QueryRow)
}
