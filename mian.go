package main

import (
	. "liuliang/dao"
	"liuliang/model"
	. "liuliang/route"
)

func main() {
	Init_db()
	DB.AutoMigrate(&model.User_infor{})
	defer DB.Close()
	r := Init()
	r.Run()

}
