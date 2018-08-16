package models

import (
	//"fmt"
	"position_postgre/db"
	//"position_postgre/tools"
	"time"
	//"log"
	"database/sql"
	"position_postgre/tools"
)

var (
	FIND_USER_SQL = `SELECT id,user_name,nick_name,avatar_url,openid FROM users WHERE openid=$1;`
)

type User struct {
	Id        int       `db:"id"`
	CreateAt  time.Time `db:"create_at"`
	UpdatedAt time.Time `db:"update_at"`
	DeleteAt  time.Time `db:"delete_at"`
	NickName  string    `db:"nick_name"`
	UserName  string    `db:"user_name"`
	Age       int       `db:"age"`
	Pwd       string    `db:"pwd"`
	Email     string    `db:"email"`
	Gender    int       `db:"gender"`
	Summary   string    `db:"summary"`
	Phone     string    `db:"phone"`
	IsDelete  int       `db:"is_delete"`
	OpenId    string    `db:"openid"`
	AvatarUrl string    `db:"avatar_url"`
}

func FindUserByOpenid(openId string) *sql.Rows {

	sqlStatement := FIND_USER_SQL
	rows, err := db.DB.Query(sqlStatement, openId)
	//fmt.Println(err)
	//defer rows.Close()

	tools.PanicError(err)
	return rows
}

func InsertUser(user User) {
	BuildInsertSql("INSERT INTO users", user, db.DB.QueryRow)
}
