package main

import (
	"github.com/gin-gonic/gin"
	"gin-practice/main/service"
)


func Router() *gin.Engine{
	r := gin.Default()
	r.GET("/index", service.GetIndex)
}