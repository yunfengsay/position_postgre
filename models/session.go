package models

import (
	jwtgo "github.com/dgrijalva/jwt-go"
	"position_postgre/db"
	myjwt "position_postgre/jwt"
	"position_postgre/modelStruct"
	"position_postgre/tools"
	"time"
)

var (
	SELECT_SQL         = `SELECT id,token,expire,create_at FROM session WHERE openid=$1`
	INSERT_SESSION_SQL = `INSERT INTO session (session, openid,expire,create_at) 
							VALUES ($1,$2,$3,$4)
							ON CONFLICT (openid) DO UPDATE 
						SET session = excluded.session, 
      openid = excluded.openid,
      expire = excluded.expire,
            create_at = excluded.create_at;`
	SecretKey = "this is secret key"
)

func MakeToken(openid string) string {
	rows, err := db.DB.Query(SELECT_SQL, openid)
	tools.PanicError(err)
	var session modelStruct.Session
	for rows.Next() {
		rows.Scan(&session.Id, &session.Session, &session.Expire, &session.CreateAt)
	}
	if session.Id != 0 {

	}
	token := tools.CreateHashWithSalt(openid)
	return token
}

func CreateSession(id int, user_name, openid string) string {
	j := &myjwt.JWT{
		[]byte("location signingkey"),
	}
	claims := myjwt.CustomClaims{
		id,
		user_name,
		openid,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),   // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 604800), // 过期时间 一个礼拜
			Issuer:    "yunfengsay",                      //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)
	tools.PanicError(err)
	return token
}

func SessionBind(openid string, session string) {
	db.DB.QueryRow(INSERT_SESSION_SQL, session, openid, time.Now().AddDate(0, 0, 7), time.Now())
}
