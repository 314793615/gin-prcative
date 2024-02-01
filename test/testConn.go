package main
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
	"gin-practice/models"
)

func main(){
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/test1?charset=utf8mb4&pasreTime=True&loc=Local"), &gorm.Config{})
	if err != nil{
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.UserBasic{})

	user:=&models.UserBasic{}
	user.Name = "张三"

	db.Create(user)

	fmt.Println(db.First(user, 1))

	db.Model(user).Update("phone", 123456789)
	db.Model(user).Update("password", 1234)

}