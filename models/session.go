package models

import (
	"github.com/satori/go.uuid"
	"position_postgre/db"
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

func CreateSession() string {
	u2, err := uuid.NewV4()
	tools.PanicError(err)
	return u2.String()
}

func SessionBind(openid string, session string) {
	db.DB.QueryRow(INSERT_SESSION_SQL, session, openid, time.Now().AddDate(0, 0, 7), time.Now())
}
