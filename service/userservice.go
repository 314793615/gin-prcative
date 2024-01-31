package service

import (
	"gin-practice/models"

	"github.com/gin-gonic/gin"
)

// FindUserByNameAndPwd
// @Summary 查询用户
// @Tag 用户模块
// @Param name query string false "用户名"
// @Param password query string false "密码"
// @Success 200 {string} json{"code": "message"}
// @Router /user/findUserByNameAndPwd [get]
func FindUserByNameAndPwd(c *gin.Context) {
	//data := make([]*models.UserBasic, 10)
	//data = models.GetUserList()
	data := models.UserBasic{}
	name := c.Query("name")
	password := c.Query("password")
	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"message": "该用户不存在",
		})
		return
	}

	flag := utils.ValidatePassword(password, user.Salt, user.PassWord)
	if !flag {
		c.JSON(200, gin.H{
			"message": "密码不正确",
		})
	}
	pwd := utils.MakePassword(password, user.Salt)
	data := models.FindUserByNameAndPwd(name, pwd)

	c.JSON(200, gin.H{
		"message": data,
	})

}
