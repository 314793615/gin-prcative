package models

import (
	"fmt"
	"gin-practice/utils"
	"time"

	"golang.org/x/text/unicode/rangetable"
	"gorm.io/gorm"
)

type UserBasic struct{
	gorm.Model
	Name string
	PassWord string
	Phone string
	Email string
	Identity string
	ClientIp string
	ClientPort string
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