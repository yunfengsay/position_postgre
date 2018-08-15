package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type GeoJson struct {
	Type        string    `json:"-"`
	Coordinates []float64 `json:"coordinates"`
}
type Location struct {
	Id         bson.ObjectId `bson:"_id" json:"id"`
	LType      []int         `bson:"l_type" json:"l_type"`
	CreateAt   time.Time     `bson:"create_at" json:"create_at"`
	UpdatedAt  time.Time     `bson:"update_at" json:"update_at"`
	DeleteAt   time.Time     `bson:"delete_at" json:"delete_at"`
	Imgs       []string      `bson:"imgs" json:"imgs"`
	Location   GeoJson       `bson:"location" json:"location"`
	Content    string        `bson:"content" json:"content"`
	User       bson.ObjectId `bson:"user" json:"user"`
	UserObj    User          `bson:"user_obj" json:"user_obj"`
	IsDelete   int           `bson:"is_delete" json:"is_delete"`
	ViewNum    int64         `bson:"viewd_num" json:"viewd_num"`
	LikedNum   int64         `bson:"liked_num" json:"liked_num"`
	CommentNum int64         `bson:"comment_num" json:"connent_num"`
	Liked      []string      `bson:"liked" json:"liked"`
}

type LocationAction struct {
	AddLocation    func(l *Location) (err error)
	UpdateLocation func(l *Location) (err error)
	DeleteLocation func(l *Location) (err error)
	GetLocation    func(id bson.ObjectId) (l Location, err error)
	GetLocations   func(id bson.ObjectId) (l Location, err error)
	NeerLocation   func()
}

type AnyLocations struct {
	Ok      int                    `json:"ok"`
	Results []interface{}          `json:"results"`
	Status  map[string]interface{} `json:"status"`
}
