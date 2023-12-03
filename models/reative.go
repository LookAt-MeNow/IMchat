package models

import (
	"ginchat/utils"

	"gorm.io/gorm"
)

//人员关系
type Reative struct {
	gorm.Model
	OwnerId 	uint	//关系所有人
	TargetId 	uint	//对应的人
	Type		int		//对应的类型
	Desc		string	//描述
}

func (table *Reative) TableName() string {
	return "reative"
}

//查找好友
func FindFriend(userId uint) []UserBasic {
	relative := make([]Reative, 0)
	utils.DB.Where("owner_id = ? and type=1", userId).Find(&relative)
	users := make([]UserBasic,0)
	for _,v := range relative {
		if v.TargetId != userId {
			user := UserBasic{}
			utils.DB.Where("id = ?", v.TargetId).First(&user)
			users = append(users, user)
		}
	}
	return users
}

//添加好友
func AddFriend(userId uint,targetId uint) (int,string){
	user := UserBasic{}
	if targetId != 0 {
		user = AddFriendByID(targetId)
		if user.Salt != "" {
			if userId == user.ID {
				return -1 , "我加我自己,人格分裂喵"
			}
			p := Reative{}
			utils.DB.Where("owner_id = ? and target_id = ? and type=1",userId,targetId).Find(&p)
			if p.ID != 0 {
				return -1, "已经是好友了喵"
			}
			//开启一个事务
			tx := utils.DB.Begin()
			defer func (){
				if r:= recover();r != nil {
					tx.Rollback()	//事务回滚
				}
			}()
			r := Reative{}
			r.OwnerId = userId
			r.TargetId = targetId
			r.Type = 1
			r.Desc = "好友"
			if err := utils.DB.Create(&r).Error;err != nil {
				tx.Rollback()	//事务回滚
				return -1,"添加失败"
			}
			q := Reative{}
			q.OwnerId = targetId
			q.TargetId = userId
			q.Type = 1
			q.Desc = "好友"
			if err := utils.DB.Create(&q).Error;err != nil {
				tx.Rollback()
				return -1,"添加失败"
			}
			tx.Commit()
			return 0,"添加成功"
		}
		return -1,"没有找到此用户"
	}
	return -1,"请输入有效字段"
}