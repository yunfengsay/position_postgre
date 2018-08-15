package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Comments struct {
	Id         bson.ObjectId `bson: "_id"`
	CreateAt   time.Time     `bson: "create_at"`
	LocationId bson.ObjectId `bson: "location_id"` // location的id
	To         bson.ObjectId `bson: "to"`          // location 的 user id
	From       bson.ObjectId `bson: "from"`        // 评论者的id
	UserAvater string        `bson: "user_avater"` // 评论者的头像
	UserName   string        `bson: "user_name"`   // 评论者的昵称
}
