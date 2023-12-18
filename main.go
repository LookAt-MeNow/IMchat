package main

import (
	"ginchat/router"
	"ginchat/utils"

	//"gorm.io/gorm"
)

//var db *gorm.DB

func main() {
	utils.InitConfig() //初始化配置文件
	utils.InitMySQL()  //连接Mysql
	utils.InitRedis()  //连接Redis
	r := app.Router()  //修改测试
	r.Run()
}