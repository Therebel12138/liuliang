package route

import (
	"github.com/gin-gonic/gin"
	ctrl "liuliang/controller"
)

func Init() *gin.Engine {
	r := gin.Default()
	usergroup := r.Group("/user")
	usergroup.POST("/register", ctrl.Userregister)
	usergroup.POST("/login", ctrl.Userlogin)
	usergroup.POST("/updatapasswd", ctrl.Usernewpasswd)
	return r
}
