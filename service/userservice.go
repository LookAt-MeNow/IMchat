package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetIndex
// @Summary 用户列表
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := models.GetUserList()

	c.JSON(200, gin.H{
		"code":	0,
		"message": data,
	})
}

// CreateUser
// @Summary 创建用户
// @Tags 用户模块
// @Param username query string false "用户名"
// @Param password query string false "密码"
// @Param repassword query string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	//user.Username = c.Query("username")
	//password := c.Query("password")
	//repassword := c.Query("repassword")

	user.Username   = c.Request.FormValue("name")
	password 		:= c.Request.FormValue("password")
	repassword 		:= c.Request.FormValue("Identity")
	//加密的随机数
	salt := fmt.Sprintf("%06d",rand.Int31())
	//判断用户是否存在
	data := models.AddFriendByUser(user.Username)
	if user.Username == "" || password == "" || repassword == ""{
		c.JSON(-1, gin.H{
			"code": 	-1,
			"message": "用户名或密码不能为空",
		})
		return
	}
	if  data.Username != "" {
		c.JSON(-1, gin.H{
			"code": 	-1,
			"message": "用户名已存在",
		})
		return
	}
	if password != repassword {
		c.JSON(-1, gin.H{
			"code": 	-1,
			"message": "两次密码不一致",
		})
		return
	}
	//user.Password = password
	//MD5密码加密
	user.Password = utils.MakePassword(password,salt)
	user.Salt = salt
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"code":	0,
		"message": "注册成功",
		"data":	user,
	})
}


// FindUserByNameAndPasswd
// @Summary 用户登录
// @Tags 用户模块
// @Param username query string false "用户名"
// @Param password query string false "密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/login [post]
func FindUserByNameAndPasswd(c *gin.Context) {
	data := models.UserBasic{}
	//name := c.Query("username")
	//password := c.Query("password")
	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	user :=	models.AddFriendByUser(name)
	if user.Username == "" {
		c.JSON(200, gin.H{
			"code":		-1,//	成功 -1 失败
			"message": 	"用户不存在",
			"data":		data,
		})
		return
	}

	flag := utils.ValidPassword(password,user.Salt,user.Password)
	if !flag {
		c.JSON(200, gin.H{
			"code":		-1,//	成功 -1 失败
			"message": 	"密码不正确",
			"data":		data,
		})
		return
	}
	pwd := utils.MakePassword(password,user.Salt)
	data = models.FindUserByNameAndPasswd(name,pwd)
	c.JSON(200, gin.H{
		"code":		0,//	成功 -1 失败
		"message": 	"登录成功",
		"data":		data,
	})
}

// DeletUser
// @Summary 删除用户
// @Tags 用户模块
// @Param id query string false "id"
// @Success 200 {string} json{"code","message"}
// @Router /user/deletUser [get]
func DeletUser(c *gin.Context) {
	user := models.UserBasic{}
	id, err := strconv.Atoi(c.Query("id"))
	if err!=nil {
		panic(err)
	}
	user.ID = uint(id)
	models.DeletUser(user)
	c.JSON(200, gin.H{
		"code":		0,//	成功 -1 失败
		"message": 	"删除成功",
		"data":		user,
	})
}

// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @Param id formData string false "id"
// @Param name formData string false "name"
// @Param password formData string false "password"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdataUser(c *gin.Context) {
	user := models.UserBasic{}
	id, err := strconv.Atoi(c.PostForm("id"))
	if err!=nil {
		panic(err)
	}
	user.ID = uint(id)
	user.Username = c.PostForm("name") //PostForm获取表单数据
	user.Password = c.PostForm("password")
	user.Avatar = fmt.Sprintf("http://localhost:8080/asset/upload/%s",FileName)
	models.UpdateUser(user)
	c.JSON(200, gin.H{
		"code":		0,//	成功 -1 失败
		"message": "修改用户成功",
	})
}

func FindByID(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Request.FormValue("userId"))

	//	name := c.Request.FormValue("name")
	data := models.AddFriendByID(uint(userId))
	utils.RepOKList(c.Writer, data, "ok")
}

func SendUserMsg(c *gin.Context){
	models.Chat(c.Writer,c.Request)
}


