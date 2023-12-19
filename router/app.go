package app

import (
	"ginchat/docs"
	"ginchat/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine{
	r := gin.Default()

	//swagger，测试接口
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swaggo/*any",ginSwagger.WrapHandler(swaggerfiles.Handler))

	//静态资源加载
	r.Static("/asset","asset/")
	r.StaticFile("/favicon.ico", "asset/images/favicon.ico")
	r.LoadHTMLGlob("statics/**/*")
	//首页
	r.GET("/",service.GetIndex)
	r.GET("/index",service.GetIndex)
	r.GET("/signup",service.SignUp)
	r.GET("/chat",service.BeginChat)
	r.GET("/initchat",service.Chat)

	//用户操作
	r.GET("/user/getUserList",service.GetUserList)
	r.POST("/user/createUser",service.CreateUser)
	r.GET("/user/deletUser",service.DeletUser)
	r.POST("/user/updateUser",service.UpdataUser)
	r.POST("/user/login",service.FindUserByNameAndPasswd)
	r.POST("/user/find",service.FindByID)

	//好友操作
	r.POST("/findFriend",service.FindFriend)       //查找好友
	r.POST("/relative/addfriend",service.AddFriend)//添加好友

	//发送消息
	r.GET("/user/sendMsg",service.SendMsg)         //发送消息
	r.GET("/user/sendUserMsg",service.SendUserMsg) //和上面的区别是，这个是给指定用户发送消息
	r.POST("/attach/upload",service.Upload)        //上传文件,音频和图片

	//群操作
	r.POST("/relative/createcommunity",service.CreateCommunity)	//创建群
	r.POST("/relative/loadcommunity",service.LoadCommunity)	//群列表
	r.POST("/relative/joincommunity",service.JoinCommunity)	//加入群
	return r
}
 