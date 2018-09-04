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
	FIND_USER_SQL        = `SELECT id,user_name,nick_name,avatar_url,openid FROM users WHERE openid=$1;`
	FIND_USER_BY_SESSION = `SELECT id,nick_name,gender,avatar_url FROM users WHERE users.openid=(
		SELECT openid FROM session WHERE session=$1
	)`
)

func FindUserByOpenid(openId string) *sql.Rows {
	rows, err := db.DB.Query(FIND_USER_SQL, openId)
	//fmt.Println(err)
	//defer rows.Close()

	tools.PanicError(err)
	return rows
}

func FindUserBySession(session string) modelStruct.User {
	rows, err := db.DB.Query(FIND_USER_BY_SESSION, session)
	tools.PanicError(err)
	var user = modelStruct.User{}
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.NickName, &user.Gender, &user.AvatarUrl)
		tools.PanicError(err)
	}
	return user
}

func InsertUser(user modelStruct.User) {
	BuildInsertSql("INSERT INTO users", user, db.DB.QueryRow)
}
