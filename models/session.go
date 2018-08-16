package models

import (
	"github.com/satori/go.uuid"

	"position_postgre/db"
	"position_postgre/tools"
	"sync"
	"time"
)

var (
	SELECT_SQL         = `SELECT id,token,expire,create_at FROM session WHERE openid=$1`
	INSERT_SESSION_SQL = `INSERT INTO session (session, openid,expire,create_at) 
							VALUES ($1,$2,$3,$4)
							ON CONFLICT (session) DO UPDATE
							`
)

type Manager struct {
	cookieName  string     // private cookiename
	lock        sync.Mutex // protects session
	provider    Provider
	maxLifeTime int64
}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

type Session struct {
	Session  string    `db:"session"`
	Id       int32     `db:"id"`
	Openid   string    `db:"openid"`
	Expire   time.Time `db:"expire"`
	CreateAt time.Time `db:"create_at"`
}

func (session *Session) SessionDestroy(sid string) {

}

func MakeToken(openid string) string {
	rows, err := db.DB.Query(SELECT_SQL, openid)
	tools.PanicError(err)
	var session Session
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
	db.DB.Query(INSERT_SESSION_SQL, session, openid, time.Hour*24*7, time.Now())
}
