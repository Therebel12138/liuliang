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
	isroot1, _ := strconv.ParseBool(isroot)
	le, _ := strconv.Atoi(secure_level)
	if orm.Findname(username) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户名重复",
		})
	} else {
		orm.Dbuserlogin(username, passwd, isroot1, le)
		c.JSON(http.StatusOK, gin.H{
			"status_mag": "注册成功",
		})
	}
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

func User_backinfor(c *gin.Context) {
	username := c.Query("username")
	if orm.Findname(username) {
		var userinfor = orm.Back_userinfor(username)
		c.JSON(http.StatusOK, userinfor)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "不存在该用户",
		})
	}
}

func Userdel(c *gin.Context) {

	username := c.PostForm("username")
	isroot := c.PostForm("isrooot")
	sir, _ := strconv.ParseBool(isroot)
	if !sir {
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户权限不够，请联系管理员",
		})
	} else {
		orm.Userdel(username)
		c.JSON(http.StatusOK, gin.H{
			"msg": "成功删除",
		})
	}
}
