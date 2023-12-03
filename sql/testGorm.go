package main

import (
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/chat?charset=utf8mb4&parseTime=True&loc=Local"))
	if err!=nil {
		panic(err)
	}
	//模型绑定
	db.AutoMigrate(&models.UserBasic{})
	db.AutoMigrate(&models.Community{})
	db.AutoMigrate(&models.GroupInfo{})
	db.AutoMigrate(&models.Reative{})
	//创建
	//user := &models.UserBasic{}
	//user.Username = "test"
	//db.Create(user)
	//查询
	//db.Model(user).Update("Password", "1234")

}