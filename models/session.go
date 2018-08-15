package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Session struct {
	Token    string        `json:"token"`
	Id       bson.ObjectId `json:"id"`
	Openid   string        `json:"openid"`
	Expire   time.Time     `json:"expire"`
	CreateAt time.Time     `json:"create_at"`
}
