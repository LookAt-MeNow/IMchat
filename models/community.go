package models

import (
	"ginchat/utils"

	"gorm.io/gorm"
)

type Community struct {
	gorm.Model
	Name 	string
	OwnerId uint
	Image  	string
	Desc 	string
}

func CreateCommunity(c Community) (int ,string){
	if len(c.Name) == 0 {	
		return -1,"群名不能为空"
	}
	if err := utils.DB.Create(&c).Error;err != nil {
		return -1,"建群失败"
	}
	return 0,"建群成功喵"
}

func LoadCommunity(ownerid uint) ([]*Community,string){
	community := make([]*Community,0)
	utils.DB.Where("owner_id = ?",ownerid).Find(&community)
	return community,"加载成功喵"
}