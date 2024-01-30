package main


import (
	"gin-practice/routers"
)

func main(){
	r := routers.Router()
	r.Run("8081")
}
