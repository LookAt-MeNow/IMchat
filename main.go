package main

import (
	"ginchat/router"
	"ginchat/utils"

	//"gorm.io/gorm"
)

//var db *gorm.DB

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	r := app.Router() //修改测试
	r.Run()
}