package apis

import (
	"encoding/json"
	"net/http"
	"position_postgre/modelStruct"
	"position_postgre/models"
	"position_postgre/tools"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	APPID  = "wx4c26683a35d3fab2"
	SECRET = "28397145e103571cfd04867d849501bf"
)

type WXUserLoginForm struct {
	Code     string   `binding:"required" json:"code"`
	UserInfo UserInfo `binding:"required" json:"user_info"`
}

type UserInfo struct {
	AvatarUrl string `binding:"required" json:"avatarUrl"`
	Gender    int    `binding:"required" json:"gender"`
	NickName  string `binding:"required" json:"nickName"`
}

type OpenID struct {
	Openid string
}

func AddUserApi(c *gin.Context) {
	pwd := c.Request.FormValue("pwd")
	nUser := modelStruct.User{}
	nUser.UserName = c.Request.FormValue("user_name")
	nUser.NickName = c.Request.FormValue("nick_name")
	nUser.Age, _ = strconv.Atoi(c.Request.FormValue("age"))
	nUser.Pwd = tools.CreateHashWithSalt(c.Request.FormValue("pwd"))
	nUser.Email = c.Request.FormValue("email")
	nUser.Gender, _ = strconv.Atoi(c.Request.FormValue("gender"))
	nUser.Summary = c.Request.FormValue("summary")
	nUser.Phone = c.Request.FormValue("phone")
	nUser.AvatarUrl = c.Request.FormValue("avatarurl")
	nUser.Pwd = pwd
	models.InsertUser(nUser)
	// err, _ := models.AddUser(nUser)
	// if err != nil {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code": 1,
	// 	})
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code": 0,
	// 	})
	// }
}

func AddWxUser(user UserInfo, nickname string) {
	nUser := modelStruct.User{}
	nUser.AvatarUrl = user.AvatarUrl
	nUser.Gender = user.Gender
	nUser.NickName = user.NickName
	nUser.OpenId = nickname
	models.InsertUser(nUser)

}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(target)
	return err
}

func WXLogin(c *gin.Context) {
	var session string
	wx_login := &WXUserLoginForm{}
	err := c.BindJSON(wx_login)
	tools.PanicError(err)
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + APPID + "&secret=" + SECRET + "&js_code=" + wx_login.Code + "&grant_type=authorization_code"
	respId := OpenID{}
	getJson(url, &respId)
	var user = modelStruct.User{}
	row := models.FindUserByOpenid(respId.Openid)
	for row.Next() {
		row.Scan(&user.Id, &user.UserName, &user.NickName, &user.AvatarUrl, &user.OpenId)
	}
	tools.PanicError(err)
	if user.Id == 0 {
		AddWxUser(wx_login.UserInfo, respId.Openid)
	}
	session = models.CreateSession()
	models.SessionBind(respId.Openid, session)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"token": session,
	})
}
