package main

import (
	"gin-practice/routers"

	"gin-practice/utils"
)

func main(){
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	r := routers.Router()
	r.Run("8082")
}
