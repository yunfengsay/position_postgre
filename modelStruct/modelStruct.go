package modelStruct

import (
	"time"
)

type Location struct {
	Id         int32     `db:"_id" json:"id"`
	LType      []int     `db:"l_type" json:"l_type"`
	CreateAt   time.Time `db:"create_at" json:"create_at"`
	UpdatedAt  time.Time `db:"update_at" json:"update_at"`
	DeleteAt   time.Time `db:"delete_at" json:"delete_at"`
	Imgs       []string  `db:"imgs" json:"imgs"`
	Location   GeoJson   `db:"location" json:"location"`
	Content    string    `db:"content" json:"content"`
	User       string    `db:"user" json:"user"`
	UserObj    User      `db:"user_obj" json:"user_obj"`
	IsDelete   int       `db:"is_delete" json:"is_delete"`
	ViewNum    int64     `db:"viewd_num" json:"viewed_num"`
	LikedNum   int64     `db:"liked_num" json:"liked_num"`
	CommentNum int64     `db:"comment_num" json:"connect_num"`
	Liked      []string  `db:"liked" json:"liked"`
}
type GeoJson struct {
	Type        string    `json:"-"`
	Coordinates []float64 `json:"coordinates"`
}
type AddLocationApiForm struct {
	Imgs    string    `binding:"required"`
	Point   []float64 `binding:"required"`
	Content string
	L_type  []int `binding:"required" json:"l_type"`
}

type AnyLocations struct {
	Ok      int                    `json:"ok"`
	Results []interface{}          `json:"results"`
	Status  map[string]interface{} `json:"status"`
}

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

type Session struct {
	Session  string    `db:"session"`
	Id       int32     `db:"id"`
	Openid   string    `db:"openid"`
	Expire   time.Time `db:"expire"`
	CreateAt time.Time `db:"create_at"`
}
