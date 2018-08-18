package models

import (
	"fmt"
	//"github.com/twpayne/go-geom"
	//"github.com/twpayne/go-geom/encoding/ewkb"
	"log"
	"position_postgre/db"
	"position_postgre/modelStruct"
)

var (
	//INSERT_LOCATION = `
	//	INSERT INTO location l_type,content,imgs,location VALUES
	//	(%d,%s,%v,ST_SetSRID(ST_MakePoint(%f,%f), 4326))
	//`
	INSERT_LOCATION = `
		INSERT INTO location (l_type,content,imgs,lat,lng) VALUES 
		(%d,'%s','%s',%f,%f)
	`
)

func AddLocation(l *modelStruct.AddLocationApiForm) {
	//rows, err := db.DB.Query(INSERT_LOCATION, l.L_type, l.Content, l.Imgs)
	insertSql := fmt.Sprintf(INSERT_LOCATION, l.L_type[0], l.Content, l.Imgs, l.Point[0], l.Point[1])
	if _, err := db.DB.Exec(insertSql); err != nil {
		log.Println(err)
	}
}
