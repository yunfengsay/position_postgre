package tools

import "log"

type LogInterface interface {
	Info(string)
	Error(interface{})
	Debug(interface{})
}

var Log LogInterface

func Info(info string) {
	log.Println(info)
}

func Error(info interface{}) {
	log.Println(info)
}
func Debug(info interface{}) {
	log.Println(info)
}

func init() {

	log.SetFlags(log.Ldate | log.Lshortfile)
}
