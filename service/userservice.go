package service

import (
	"gin-practice/models"

	"github.com/gin-gonic/gin"
	"github.com/pelletier/go-toml/v2/internal/danger"
)

// GetUserList
// @Tag 首页
// @Success 200 {string} json{"code", "message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context){
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})

}

// CreateUser
// @Tags 用户模块
// @parm name query string false "用户名"
// @parm password query string false "密码"
// @parm repassword query string false "确认密码"
// @Success 200 {string} json{"code", "message"}
// @Router /user/createUser [Get]
func CreateUser(c *gin.Context){
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	if password != repassword{
		c.JSON(-1, gin.H{
			"message": "两次密码不一致",
		})
	}
	user.PassWord = password
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"message":"新增用户成功",
	})
}