package models

import "gorm.io/gorm"

//群
type GroupInfo struct {
	gorm.Model
	Name 		string
	OwnerId 	uint
	Icon		string
	Type		int
	Desc 		string
}

func (table *GroupInfo) TableName() string {
	return "group_info"
}