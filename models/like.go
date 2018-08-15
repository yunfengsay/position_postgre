package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Like struct {
	Id       bson.ObjectId `bson: "_id"`
	CreateAt time.Time     `bson: "create_at"`
	Location bson.ObjectId `bson: "location"` // location的id
	To       bson.ObjectId `bson: "to"`       // location 的 user id
	From     bson.ObjectId `bson: "from"`     // 点击者的id
}

type LikeAction struct {
	AddLike    func(l *Like) (err error)
	DeleteLike func(id bson.ObjectId) (err error)
}
