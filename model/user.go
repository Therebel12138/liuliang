package model

import (
	"liuliang/dao"
)

type User_infor struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Passwd       string `json:"passwd"`
	Isroot       bool   `json:"isroot"`
	Secure_level int    `json:"secure_level"`
}

func Dbuserlogin(name string, passwd string, isroot bool, scec int) {
	user := User_infor{Name: name, Passwd: passwd, Isroot: isroot, Secure_level: scec}
	dao.DB.Create(&user)
}

func Userlogin(username string, passwd string) bool {
	var user User_infor
	dao.DB.Where("name = ? AND passwd = ?", username, passwd).First(&user)
	if user.Id != 0 {
		return true
	} else {
		return false
	}
}

func Uppasswd(username string, newpasswd string) bool {
	var user User_infor
	user.Name = username
	dao.DB.Model(&user).Where("name = ?", username).Update("passwd", newpasswd)
	if user.Passwd == newpasswd {
		return true
	} else {
		return false
	}
}

func Back_userinfor(username string) User_infor {
	var user User_infor
	dao.DB.Where("name = ?", username).First(&user)
	return user
}
