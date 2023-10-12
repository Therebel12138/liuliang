package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	orm "liuliang/model"
	"net/http"
	"strconv"
	"time"
)

func Pptupload(c *gin.Context) {
	name := c.PostForm("pname")
	secure := c.PostForm("secure_level")
	updata_time := c.PostForm("update_time")
	ismaintianed := c.PostForm("ismaintained")
	asset_class := c.PostForm("asset_class")
	brand_name := c.PostForm("brand_name")
	work_dpt := c.PostForm("work_dpt")
	se, err := strconv.Atoi(secure)
	ism, err := strconv.ParseBool(ismaintianed)
	ass, err := strconv.Atoi(asset_class)
	t, err := time.ParseInLocation("2006-01-02 15:04:05", updata_time, time.Local)

	ppt := orm.Property_infor{Name: name, Secure_level: se, UpdateTime: t, Ismaintained: ism, Asset_class: ass, Brand_name: brand_name, Work_dpt: work_dpt}
	if err == nil {
		orm.Pro_upload(ppt)
	}
	c.JSON(http.StatusOK, ppt)
}

func Pptsearch(c *gin.Context) {
	searchname := c.PostForm("pname")
	secure_level := c.PostForm("secure_level")
	serc, err := strconv.Atoi(secure_level)
	if err == nil {
	}
	var ppt = orm.Pptsearch(searchname, serc)
	c.JSON(http.StatusOK, ppt)
}

func Pptupdata(c *gin.Context) {
	name := c.PostForm("pname")
	secure := c.PostForm("secure_level")
	updata_time := c.PostForm("update_time")
	ismaintianed := c.PostForm("ismaintained")
	asset_class := c.PostForm("asset_class")
	brand_name := c.PostForm("brand_name")
	work_dpt := c.PostForm("work_dpt")
	se, err := strconv.Atoi(secure)
	ism, err := strconv.ParseBool(ismaintianed)
	ass, err := strconv.Atoi(asset_class)
	t, err := time.ParseInLocation("2006-01-02 15:04:05", updata_time, time.Local)
	ppt := orm.Property_infor{Name: name, Secure_level: se, UpdateTime: t, Ismaintained: ism, Asset_class: ass, Brand_name: brand_name, Work_dpt: work_dpt}
	if err == nil {
	}
	fmt.Println(ppt.Name)
	orm.Pptupdata(ppt)
	c.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}
