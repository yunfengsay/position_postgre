package modelStruct

import (
	"encoding/json"
	"time"

	"position_postgre/tools"
)

type Location struct {
	Id         int32     `db:"id" json:"id"`
	LType      string    `db:"l_type" json:"l_type"`
	CreateAt   time.Time `db:"create_at" json:"create_at"`
	UpdatedAt  time.Time `db:"update_at" json:"update_at"`
	DeleteAt   time.Time `db:"delete_at" json:"delete_at"`
	Imgs       string    `db:"imgs" json:"imgs"`
	Point      GeoJson   `db:"point" json:"point"`
	Content    string    `db:"content" json:"content"`
	User       string    `db:"user" json:"user"`
	UserObj    string    `db:"user_obj" json:"user_obj"`
	IsDelete   int       `db:"is_delete" json:"is_delete"`
	ViewNum    int64     `db:"viewd_num" json:"viewed_num"`
	LikedNum   int64     `db:"liked_num" json:"liked_num"`
	CommentNum int64     `db:"comment_num" json:"connect_num"`
	Liked      []string  `db:"liked" json:"liked"`
}

type LocationResponse struct {
	Location Location
	Distance float32 `db:"distance" json:"distance`
}

type GeoJson struct {
	Type        string    `json:"-"`
	Coordinates []float64 `json:"coordinates"`
}
type AddLocationApiForm struct {
	Imgs    string    `binding:"required"`
	Point   []float64 `binding:"required"`
	Content string
	L_type  string `binding:"required" json:"l_type"`
}

type AnyLocations struct {
	Ok      int                    `json:"ok"`
	Results []interface{}          `json:"results"`
	Status  map[string]interface{} `json:"status"`
}

type User struct {
	Id        int       `db:"id" json:"id"`
	CreateAt  time.Time `db:"create_at" json:"create_at"`
	UpdatedAt time.Time `db:"update_at" json:"updated_at"`
	DeleteAt  time.Time `db:"delete_at" json:"delete_at"`
	NickName  string    `db:"nick_name" json:"nick_name"`
	UserName  string    `db:"user_name" json:"user_name"`
	Age       int       `db:"age" json:"age"`
	Pwd       string    `db:"pwd" json:"pwd"`
	Email     string    `db:"email" json:"email"`
	Gender    int       `db:"gender" json:"gender"`
	Summary   string    `db:"summary" json:"summary"`
	Phone     string    `db:"phone" json:"phone"`
	IsDelete  int       `db:"is_delete" json:"is_delete"`
	OpenId    string    `db:"openid" json:"open_id"`
	AvatarUrl string    `db:"avatar_url" json:"avatar_url"`
}

func (user *User) String() string {
	user_string, err := json.Marshal(user)
	tools.PanicError(err)
	return string(user_string)
}

type Session struct {
	Session  string    `db:"session"`
	Id       int32     `db:"id"`
	Openid   string    `db:"openid"`
	Expire   time.Time `db:"expire"`
	CreateAt time.Time `db:"create_at"`
}

// 获取附近的人的接口数据
type LocationRequest struct {
	LType []int     `binding:"required" json:"l_type"`
	Point []float32 `binding:"required" json:"point"`
	R     int       `binding:"required" json:"r"` // 半径范围
}
