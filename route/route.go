package route

import (
	"github.com/gin-gonic/gin"
	ctrl "liuliang/controller"
	mid "liuliang/midware"
)

func Init() *gin.Engine {
	r := gin.Default()
	usergroup := r.Group("/user")
	usergroup.POST("/register", ctrl.Userregister)
	usergroup.POST("/login", mid.Sendtoken(), ctrl.Userlogin)
	usergroup.POST("/updatapasswd", mid.Twjparse(), ctrl.Usernewpasswd)
	usergroup.GET("/back_infor", mid.Twjparse(), ctrl.User_backinfor)
	usergroup.POST("/del", mid.Twjparse(), ctrl.Userdel)
	pptgroup := r.Group("/property", mid.Twjparse())
	pptgroup.POST("/upload", ctrl.Pptupload)
	pptgroup.POST("/search", ctrl.Pptsearch)
	pptgroup.POST("/updata", ctrl.Pptupdata)
	pptgroup.POST("/upexc", ctrl.Upexcel)
	pptgroup.POST("del", ctrl.Pptdel)
	pptgroup.POST("dwexcel", ctrl.Dwexcel)
	return r
}
