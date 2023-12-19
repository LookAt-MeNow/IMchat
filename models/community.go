package models

import (
	"ginchat/utils"

	"gorm.io/gorm"
	"fmt"
)

type Community struct {
	gorm.Model
	Name 	string
	OwnerId uint
	Image  	string
	Desc 	string
}

func CreateCommunity(c Community) (int ,string){
	tx := utils.DB.Begin()
	//事务一旦开始，不论什么异常最终都会 Rollback
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if len(c.Name) == 0 {
		return -1, "群名称不能为空"
	}
	if c.OwnerId == 0 {
		return -1, "请先登录"
	}
	if err := utils.DB.Create(&c).Error; err != nil {
		fmt.Println(err)
		tx.Rollback() //事务回滚
		return -1, "建群失败"
	}
	r := Reative{}
	r.OwnerId = c.OwnerId
	r.TargetId = c.ID
	r.Type = 2 //群关系
	r.Desc = "群主"
	if err := utils.DB.Create(&r).Error; err != nil {
		tx.Rollback()
		return -1, "添加群关系失败"
	}

	tx.Commit() //事务提交
	return 0, "建群成功"
}

//加载群
func LoadCommunity(ownerid uint) ([]*Community,string){
	r := make([]Reative, 0)
	objIds := make([]uint64, 0)
	utils.DB.Where("owner_id = ? and type=2", ownerid).Find(&r)
	for _, v := range r {
		objIds = append(objIds, uint64(v.TargetId))
	}

	data := make([]*Community, 10)
	utils.DB.Where("id in ?", objIds).Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	//utils.DB.Where()
	return data, "查询成功"
}