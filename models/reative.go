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
func AddFriend(userId uint,targetName string) (int,string){
	//user := UserBasic{}
	if targetName != "" {
		targetUser := AddFriendByUser(targetName)
		if targetUser.Salt != "" {
			if userId == targetUser.ID {
				return -1 , "我加我自己,人格分裂喵"
			}
			p := Reative{}
			utils.DB.Where("owner_id = ? and target_id = ? and type=1",userId,targetUser.ID).Find(&p)
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
			r.TargetId = targetUser.ID
			r.Type = 1
			r.Desc = "好友"
			if err := utils.DB.Create(&r).Error;err != nil {
				tx.Rollback()	//事务回滚
				return -1,"添加失败"
			}
			q := Reative{}
			q.OwnerId = targetUser.ID
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

// 加群
func JoinGroup(userId uint, comId string) (int, string) {
	r := Reative{}
	r.OwnerId = userId
	//contact.TargetId = comId
	r.Type = 2
	r.Desc = "群友"
	community := Community{}

	utils.DB.Where("id=? or name=?", comId, comId).Find(&community)
	if community.Name == "" {
		return -1, "没有找到群"
	}
	utils.DB.Where("owner_id=? and target_id=? and type =2 ", userId, comId).Find(&r)
	if !r.CreatedAt.IsZero() {
		return -1, "已加过此群"
	} else {
		r.TargetId = community.ID
		utils.DB.Create(&r)
		return 0, "加群成功"
	}
}

//获取群友的ID
func GetCroupUser(communityID uint) []uint{
	r := make([]Reative, 0)	//群关系
	obj := make([]uint,0)   //群友ID
	utils.DB.Where("target_id = ? and type=2", communityID).Find(&r)
	for _, v := range r {
		obj = append(obj, uint(v.OwnerId))
	}
	return obj
}