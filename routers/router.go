package routers

import (
	"github.com/gin-gonic/gin"
	"gin-practice/service"
)


func Router() *gin.Engine{
	r := gin.Default()
	r.GET("/index", service.GetIndex)
	r.GET("/user/getUserList", service.GetUserList)
	return r
}