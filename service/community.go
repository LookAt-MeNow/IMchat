package service

import (
	"ginchat/models"
	"strconv"
	"ginchat/utils"
	"github.com/gin-gonic/gin"
)

//新建群
func CreateCommunity(c *gin.Context) {
	ownerid ,err := strconv.Atoi(c.Request.FormValue("ownerId"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "ownerId error",
		})
	}
	name  := c.Request.FormValue("name")	//字段一定要和前端一致
	community := models.Community{}
	community.OwnerId = uint(ownerid)
	community.Name 	  = name
	code ,msg := models.CreateCommunity(community)	
	if code == 0 {
		utils.RepOK(c.Writer,code,msg)
	}else {
		utils.RepFail(c.Writer,msg)
	}
}

func LoadCommunity(c *gin.Context) {
	ownerid ,err := strconv.Atoi(c.Request.FormValue("ownerId"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "ownerId error",
		})
	}
	code ,msg := models.LoadCommunity(uint(ownerid))	
	if len(code) != 0 {
		utils.RepOKList(c.Writer,code,msg)
	}else {
		utils.RepFail(c.Writer,msg)
	}
}

func JoinCommunity(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Request.FormValue("userId"))
	comId := c.Request.FormValue("comId")

	//	name := c.Request.FormValue("name")
	code, msg := models.JoinGroup(uint(userId), comId)
	if code == 0 {
		utils.RepOK(c.Writer,code,msg)
	}else {
		utils.RepFail(c.Writer,msg)
	}
}
