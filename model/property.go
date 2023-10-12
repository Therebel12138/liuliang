package model

import (
	"database/sql/driver"
	"fmt"
	"liuliang/dao"
	"time"
)

type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type Rroperty_infor struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Secure_level int       `json:"secure_level"`
	UpdateTime   time.Time `gorm:"column:updatetime;type:datetime(0);autoUpdateTime" json:"updatetime"`
	Ismaintained bool      `json:"ismaintained" `
	Asset_class  int       `json:"asset_class"`
	Brand_name   string    `json:"brand_name"`
	Work_dpt     string    `json:"work_dpt"`
}

func Pro_upload(pro Rroperty_infor) {
	dao.DB.Create(&pro)
}
