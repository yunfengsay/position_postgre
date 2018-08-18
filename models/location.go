package models

import (
	//"time"
	"position_postgre/modelStruct"
)

func AddLocation(l *modelStruct.AddLocationApiForm) {
	var location = modelStruct.Location{}
	location.LType = l.L_type
	//location.Location = l.Point
	location.Content = l.Content
	location.Imgs = l.Imgs

}
