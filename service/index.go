package service

import (
	"fmt"
	"ginchat/models"
	"html/template"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetIndex
// @Tags 首页
// @Success 200 {string} Helloworld
// @Router /index [get]
func GetIndex(c *gin.Context) {
	//c.JSON(200, gin.H{
	//	"message": "hello world",
	//})
	index,err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	index.Execute(c.Writer,nil)

}

func SignUp(c *gin.Context) {
	//c.JSON(200, gin.H{
	//	"message": "hello world",
	//})
	index,err := template.ParseFiles("statics/user/register.html")
	if err != nil {
		panic(err)
	}
	index.Execute(c.Writer,nil)
}

func BeginChat(c *gin.Context) {
	//c.JSON(200, gin.H{
	//	"message": "hello world",
	//})
	index, err := template.ParseFiles("statics/chat/index.html",
		"statics/chat/head.html",
		"statics/chat/foot.html",
		"statics/chat/tabmenu.html",
		"statics/chat/concat.html",
		"statics/chat/group.html",
		"statics/chat/profile.html",
		"statics/chat/createcom.html",
		"statics/chat/userinfo.html",
		"statics/chat/main.html")
	if err != nil {
		panic(err)
	}
	userid,_ :=strconv.Atoi(c.Query("userId")) 
	token  :=c.Query("token")
	user   := models.UserBasic{}
	user.ID = uint(userid)
	user.Identity  = token
	fmt.Println("user>>>>>>>>>>>>>:",user ,"||", user.Identity)
	index.Execute(c.Writer,user)
}

