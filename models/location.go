package models

import (
	"fmt"
	//"github.com/twpayne/go-geom"
	//"github.com/twpayne/go-geom/encoding/ewkb"
	"log"
	"position_postgre/db"
	"position_postgre/modelStruct"
	"position_postgre/tools"
)

var (
	//INSERT_LOCATION = `
	//	INSERT INTO location l_type,content,imgs,location VALUES
	//	(%d,%s,%v,ST_SetSRID(ST_MakePoint(%f,%f), 4326))
	//`
	INSERT_LOCATION = `
		INSERT INTO location (l_type,content,imgs,point,lat,lng,user_obj) VALUES 
		($1,$2,$3, ST_GeomFromText('POINT(%f %f)',4326),%f,%f,$4)
	`
	NEER_LOCATION = `
	SELECT id,imgs,content,create_at,l_type,user_obj,comment_num,liked_num,ST_Distance_Sphere(Point(%f, %f)::geometry ,point) AS distance FROM location ORDER BY distance limit 10;
`
)

//func makeInserSql(l_type int, content string, imgs string, lat float32, lng float32) string {
//	return fmt.Sprintf(INSERT_LOCATION, l_type, content, imgs, lat, lng)
//}

func AddLocation(l *modelStruct.AddLocationApiForm, user modelStruct.User) {
	//rows, err := db.DB.Query(INSERT_LOCATION, l.L_type, l.Content, l.Imgs)
	insert_location_with_point := fmt.Sprintf(INSERT_LOCATION, l.Point[0], l.Point[1], l.Point[0], l.Point[1])
	if _, err := db.DB.Exec(insert_location_with_point, l.L_type, l.Content, l.Imgs, user.String()); err != nil {
		log.Println(err)
	}
}

func NeerLocation(request *modelStruct.LocationRequest) []modelStruct.LocationResponse {
	insertSql := fmt.Sprintf(NEER_LOCATION, request.Point[0], request.Point[1])
	rows, err := db.DB.Query(insertSql)
	defer rows.Close()
	location := modelStruct.LocationResponse{}
	locations := []modelStruct.LocationResponse{}
	tools.PanicError(err)
	for rows.Next() {
		err := rows.Scan(&location.Location.Id, &location.Location.Imgs, &location.Location.Content, &location.Location.CreateAt, &location.Location.LType, &location.Location.UserObj, &location.Location.CommentNum, &location.Location.LikedNum, &location.Distance)
		tools.PanicError(err)
		locations = append(locations, location)
	}
	fmt.Println(locations)
	return locations
}
