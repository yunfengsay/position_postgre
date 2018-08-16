package models

import (
	"database/sql"
	"reflect"
	"strconv"

	"log"
)

func BuildInsertSql(sqlBefore string, structType interface{}, sqlFunc func(query string, args ...interface{}) *sql.Row) {
	// example sqlBefore  `INSERT INTO USER`
	t := reflect.TypeOf(structType)
	v := reflect.ValueOf(structType)
	sqlBefore += " ("
	end := ")"
	values_after := " VALUES("

	values := []interface{}{}
	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)
		// Get the field tag value
		tag := field.Tag.Get("db")
		if tag == "id" {
			continue
		}
		sqlBefore += tag + ","
		values = append(values, v.Field(i).Interface())
		values_after += "$" + strconv.Itoa(i) + ","
		//log.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
	}
	// make sql string
	sqlBefore = sqlBefore[:len(sqlBefore)-1]
	values_after = values_after[:len(values_after)-1]
	sqlBefore += end
	values_after += end
	sql_str := sqlBefore + values_after

	// exec sql insert
	row := sqlFunc(sql_str, values...)
	log.Println(row)
}
