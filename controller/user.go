package controller

import (
	"github.com/gin-gonic/gin"
	orm "liuliang/model"
	"net/http"
	"strconv"
)

func Userregister(c *gin.Context) {
	username := c.PostForm("username")
	passwd := c.PostForm("passwd")
	isroot := c.PostForm("isroot")
	secure_level := c.PostForm("secure_level")
	isroot1, err := strconv.ParseBool(isroot)
	le, err2 := strconv.Atoi(secure_level)
	if err == nil {
		if err2 == nil {
			orm.Dbuserlogin(username, passwd, isroot1, le)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status_mag": "注册成功",
	})
}

func Userlogin(c *gin.Context) {
	username := c.PostForm("username")
	passwd := c.PostForm("passwd")
	if orm.Userlogin(username, passwd) {
		c.JSON(http.StatusOK, gin.H{
			"status_msg": "登陆成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status_msg": "密码错误",
		})
	}
}

func Usernewpasswd(c *gin.Context) {
	username := c.PostForm("username")
	passwd := c.PostForm("passwd")
	if orm.Uppasswd(username, passwd) {
		c.JSON(http.StatusOK, gin.H{
			"status_msg": "更改成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status_msg": "更新失败",
		})
	}
}
