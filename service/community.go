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
	data ,msg := models.LoadCommunity(uint(ownerid))	
	if len(data) != 0 {
		utils.RepOKList(c.Writer,data,msg)
	}else {
		utils.RepFail(c.Writer,msg)
	}
}