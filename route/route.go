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
	usergroup.GET("/back_infor", ctrl.User_backinfor)
	pptgroup := r.Group("/property")
	pptgroup.POST("/upload", ctrl.Pptupload)
	pptgroup.POST("/search", ctrl.Pptsearch)
	pptgroup.POST("/updata", ctrl.Pptupdata)
	return r
}
