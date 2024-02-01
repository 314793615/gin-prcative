package service

import (
	"fmt"
	"gin-practice/models"
	"math/rand"
	"strconv"
	"gin-practice/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
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
	user := models.FindUserByName(name)

}

// CreateUser
// @Tags 用户模块
// @Param name query string false "用户名"
// @Param password query string false "密码"
// @Param repassword query string false "确认密码"
// @Success 200 {string} json{"code", "message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context){
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")


	if password != repassword{
		c.JSON(-1, gin.H{
			"message": "两次密码不一致",
		})
		return
	}

	data := models.FindUserByName(user.Name)
	if data.Name != ""{
		c.JSON(500, gin.H{
			"Message": "用户名已注册",
		})
		return 
	}

	salt := fmt.Sprintf("%06d", rand.Int31())

	user.PassWord = utils.MakePassword(password, salt)
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"message":"新增用户成功",
	})
}

// DeleteUser
// @Tags 用户模块
// @Param ID query string false "ID"
// @Success 200 {string} json{"code", "message"} 
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context){
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"message":"删除用户成功",
	})
}

// UpdateUser
// @Tags 用户模块
// @Param id formData string false "id"
// @Param name formData string false "用户名"
// @Param password formData string false "密码"
// @Success 200 {string} json{"code", "message"}
// @Router /user/createUser [post]
func UpdateUser(c *gin.Context){
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("name"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	
	_, err := govalidator.ValidateStruct(user)
	if err != nil{
		fmt.Println(err)
		c.JSON(400, gin.H{
			"message": "修改参数不匹配",
		})
		return
	}
	models.UpdateUser(user)
	c.JSON(200, gin.H{
		"message":"修改用户成功",
	})
}

