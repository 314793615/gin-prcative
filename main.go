package main

import (
	"gin-practice/routers"

	"gin-practice/utils"
)

func main(){
	utils.InitConfig()
	utils.InitMySQL()
	r := routers.Router()
	r.Run()
}
