package models

import (
	"fmt"
	"position_postgre/db"
	//"position_postgre/tools"

	"time"
	//"log"
)

type User struct {
	Id        int       `bson:"id"`
	CreateAt  time.Time `bson:"create_at"`
	UpdatedAt time.Time `bson:"update_at"`
	DeleteAt  time.Time `bson:"delete_at"`
	NickName  string    `bson:"nick_name"`
	UserName  string    `bson:"user_name"`
	Age       int       `bson:"age"`
	Pwd       string    `bson:"pwd"`
	Email     string    `bson:"email"`
	Gender    int       `bson:"gender"`
	Summary   string    `bson:"summary"`
	Phone     string    `bson:"phone"`
	IsDelete  int       `bson:"is_delete"`
	OpenId    string    `bson:"openid"`
	AvatarUrl string    `bson:"avatar_url"`
}

func FindUserByOpenid(openId string) int {
	sqlStatement := `
		INSERT INTO users(age, email)
		VALUES ($1, '$2')
		RETURNING id`
	id := 0
	row := db.DB.QueryRow(sqlStatement, 30, "jon@calhoun.io")
	fmt.Println(row)
	fmt.Println("New record ID is:", id)
	return id
}
