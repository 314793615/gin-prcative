package models

import (
	"fmt"
	"gin-practice/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct{
	gorm.Model
	Name string 
	PassWord string
	Phone string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email string `valid:"email"`
	Identity string
	ClientIp string
	ClientPort string
	Salt string
	LoginTime time.Time
	heatBeatTime time.Time
	LogOutTime time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	isLoginOut bool
	DeviceInfo string
}

func (table *UserBasic) TableName() string{
	return "user_basic"
}


func GetUserList() []*UserBasic{
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data{
		fmt.Println(v)
	}
	return data
}

func CreateUser(user UserBasic) *gorm.DB{
	return utils.DB.Create(&user)
}

func DeleteUser(user UserBasic) *gorm.DB{
	return utils.DB.Delete(&user)
}

func UpdateUser(user UserBasic) *gorm.DB{
	return utils.DB.Model(&user).Updates(UserBasic{Name: user.Name, PassWord: user.PassWord, Phone: user.Phone, Email: user.Email})
}

func FindUserByNameAndPwd(name string, passwd string) UserBasic{
	user := UserBasic{}
	utils.DB.Where("name = ? and pass_word= ? ", name , passwd).First(&user)
	return user
}

func FindUserByName(name string) UserBasic{
	user := UserBasic{}
	utils.DB.Where("name = ?", name).First(&user)
	return user
}

func FindUserByPhone(phone string) *gorm.DB{
	user := UserBasic{}
	return  utils.DB.Where("phone = ?", phone).First(&user)
}

func FindUserByEmail(email string) *gorm.DB{
	user := UserBasic{}
	return  utils.DB.Where("email = ?", email).First(&user)
}