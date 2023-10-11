package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Init_db() {
	var err error
	DB, err = gorm.Open("mysql", "root:5buzaiaita@/liu?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}
