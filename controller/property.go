package controller

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	orm "liuliang/model"
	"net/http"
	"os"
	"path"
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
	se, _ := strconv.Atoi(secure)
	ism, _ := strconv.ParseBool(ismaintianed)
	ass, _ := strconv.Atoi(asset_class)
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", updata_time, time.Local)
	ppt := orm.Property_infor{Name: name, Secure_level: se, UpdateTime: t, Ismaintained: ism, Asset_class: ass, Brand_name: brand_name, Work_dpt: work_dpt}
	orm.Pro_upload(ppt)
	c.JSON(http.StatusOK, ppt)
}

func Pptsearch(c *gin.Context) {
	searchname := c.PostForm("pname")
	secure_level := c.PostForm("secure_level")
	serc, _ := strconv.Atoi(secure_level)
	var ppts []orm.Property_infor
	ppts = orm.Pptsearch(searchname, serc)
	c.JSON(http.StatusOK, ppts)
}

func Pptupdata(c *gin.Context) {
	name := c.PostForm("pname")
	secure := c.PostForm("secure_level")
	updata_time := c.PostForm("update_time")
	ismaintianed := c.PostForm("ismaintained")
	asset_class := c.PostForm("asset_class")
	brand_name := c.PostForm("brand_name")
	work_dpt := c.PostForm("work_dpt")
	se, _ := strconv.Atoi(secure)
	ism, _ := strconv.ParseBool(ismaintianed)
	ass, _ := strconv.Atoi(asset_class)
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", updata_time, time.Local)
	ppt := orm.Property_infor{Name: name, Secure_level: se, UpdateTime: t, Ismaintained: ism, Asset_class: ass, Brand_name: brand_name, Work_dpt: work_dpt}
	fmt.Println(ppt.Name)
	orm.Pptupdata(ppt)
	c.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}

func Upexcel(c *gin.Context) {
	files, _ := c.FormFile("pptexcel")
	dst := path.Join("D:/gowork/src/liuliang/static", files.Filename)
	newfile := c.SaveUploadedFile(files, dst)
	if newfile == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "上传成功",
		})
	}
	xlsx, _ := excelize.OpenFile(dst)
	rows := xlsx.GetRows("Sheet1")
	for key, row := range rows {
		if key > 0 {
			se, _ := strconv.Atoi(row[1])
			t, _ := time.ParseInLocation("2006-01-02 15:04:05", row[2], time.Local)
			ism, _ := strconv.ParseBool(row[3])
			a, _ := strconv.Atoi(row[4])
			ppt := orm.Property_infor{Name: row[0], Secure_level: se, UpdateTime: t, Ismaintained: ism, Asset_class: a, Brand_name: row[5], Work_dpt: row[6]}
			orm.Pro_upload(ppt)
		}
	}
	fmt.Println(files.Filename)
	os.Remove("D:/gowork/src/liuliang/static/" + files.Filename)
}

func Pptdel(c *gin.Context) {
	id := c.PostForm("ppt_id")
	secure_level := c.PostForm("secure_level")
	se, _ := strconv.Atoi(secure_level)
	id1, _ := strconv.Atoi(id)
	if orm.Pptdel(id1, se) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "删除成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "没有权限",
		})
	}

}

func Dwexcel(c *gin.Context) {
	ppt_name := c.PostForm("ppt_name")
	secure := c.PostForm("secure_level")
	sec, _ := strconv.Atoi(secure)
	pps := orm.Pptsearch(ppt_name, sec)
	xlsx := excelize.NewFile()
	index := xlsx.NewSheet("Sheet1")
	xlsx.SetCellValue("Sheet1", "A1", "name")
	xlsx.SetCellValue("Sheet1", "B1", "secure_level")
	xlsx.SetCellValue("Sheet1", "C1", "updatetime")
	xlsx.SetCellValue("Sheet1", "D1", "ismaintained")
	xlsx.SetCellValue("Sheet1", "E1", "asset_class")
	xlsx.SetCellValue("Sheet1", "F1", "brand_name")
	xlsx.SetCellValue("Sheet1", "G1", "work_dpt")
	for ind, value := range pps {
		now := strconv.Itoa(ind + 2)
		t := value.UpdateTime.Format("2006-01-02 15:04:05")
		xlsx.SetCellValue("Sheet1", "A"+now, value.Name)
		xlsx.SetCellValue("Sheet1", "B"+now, value.Secure_level)
		xlsx.SetCellValue("Sheet1", "C"+now, t)
		xlsx.SetCellValue("Sheet1", "D"+now, value.Ismaintained)
		xlsx.SetCellValue("Sheet1", "E"+now, value.Asset_class)
		xlsx.SetCellValue("Sheet1", "F"+now, value.Brand_name)
		xlsx.SetCellValue("Sheet1", "G"+now, value.Work_dpt)
	}
	xlsx.SetActiveSheet(index)
	xlsx.SaveAs("D:/gowork/src/liuliang/static/new.xlsx")
	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"new.xlsx") // 用来指定下载下来的文件名
	c.Header("Content-Transfer-Encoding", "binary")
	c.File("D:/gowork/src/liuliang/static/new.xlsx")
	os.Remove("D:/gowork/src/liuliang/static/new.xlsx")
}
