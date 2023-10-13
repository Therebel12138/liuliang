package model

import (
	"fmt"
	"liuliang/dao"
	"time"
)

type Property_infor struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Secure_level int       `json:"secure_level"`
	UpdateTime   time.Time `gorm:"column:updatetime;type:datetime(0);autoUpdateTime" json:"updatetime"`
	Ismaintained bool      `json:"ismaintained" `
	Asset_class  int       `json:"asset_class"`
	Brand_name   string    `json:"brand_name"`
	Work_dpt     string    `json:"work_dpt"`
}

func Pro_upload(pro Property_infor) {
	dao.DB.Create(&pro)
}

func Pptsearch(name string, sec int) []Property_infor {
	var ppts []Property_infor
	dao.DB.Where("name LIKE ? AND secure_level <= ?", "%"+name+"%", sec).Find(&ppts)
	return ppts
}

func Pptupdata(ppt Property_infor) {
	fmt.Println(ppt.Name)
	dao.DB.Model(Property_infor{}).Where("name = ?", ppt.Name).Updates(&ppt)
}

func Pptdel(id int, sec int) bool {
	var ppt = Property_infor{}
	dao.DB.Where("id = ?", id).First(&ppt)
	if sec >= ppt.Secure_level {
		dao.DB.Where("id = ?", id).Delete(Property_infor{})
		return true
	} else {
		return false
	}
}
