package models

import (
	"fmt"
	"ginchat/utils"
	"time"
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Username 	string 
	Password 	string 
	Phone 		string 
	Email 		string 
	Identity 	string 
	ClientIp 	string 
	ClirntPort 	string 
	Salt 		string
	Avatar      string
	Logintime 	uint64 
	Hearbeat 	uint64 
	Logouttime 	uint64 
	IsLogout 	bool 	
	DeviceInfo 	string
}

func (table *UserBasic)	TableName() string {
	return "user_basic"
}

//用户列表
func GetUserList() []*UserBasic{
	data := make([] *UserBasic,10)
	utils.DB.Find(&data)
	for _,v := range data {
		fmt.Println(v)
	}
	return data
}

//创建用户
func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

//删除用户 
func DeletUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}

//修改用户
func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Updates(UserBasic{Username: user.Username, Password: user.Password, Avatar: user.Avatar})
}

//查找用户by用户名
func AddFriendByUser(name string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("Username = ?",name).First(&user)
	return user
}

//查找用户by手机号
func AddFriendByPhone(phone string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("Phone = ?",phone).First(&user)
}

//查找用户by邮箱
func AddFriendEmail(email string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("Email = ?",email).First(&user)
}
//查找用户byID
func AddFriendByID(id uint) UserBasic {
	user := UserBasic{}
	utils.DB.Where("id = ?",id).First(&user)
	return user
}
//登录
func FindUserByNameAndPasswd(name string,passwd string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("Username = ? and Password =?",name,passwd).First(&user)

	//token加密
	str := fmt.Sprintf("%d",time.Now().Unix())
	temp := utils.MD5Code(str)
	utils.DB.Model(&user).Where("id = ?",user.ID).Update("Identity",temp)
	return user
}