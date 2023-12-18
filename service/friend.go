package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindFriend(c *gin.Context) {
	userIdStr := c.Request.FormValue("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	users := models.FindFriend(uint(userId)) 
	if err != nil {
		c.JSON(-1, gin.H{
			"code":    -1,
			"message": "参数错误",
		})
		return
	}
	utils.RepOKList(c.Writer, users, len(users))
}	


func AddFriend(c *gin.Context) {
	userId, err := strconv.Atoi(c.Request.FormValue("userId")) 
	if err != nil {
		c.JSON(-1, gin.H{
			"code":    -1,
			"message": "userId参数错误",
		})
		return
	}
	targetName := c.Request.FormValue("targetName")
/* 	targetId, err := strconv.Atoi(c.Request.FormValue("targetId"))
	if err != nil {
		c.JSON(-1, gin.H{
			"code":    -1,
			"message": "targetId参数错误",
		})
		return
	} */
	fmt.Println("1.",userId," 2.",targetName)
	code,msg := models.AddFriend(uint(userId), targetName)
	if code == 0 {
		utils.RepOK(c.Writer,code,msg)
	}else {
		utils.RepFail(c.Writer,msg)
	}
}
