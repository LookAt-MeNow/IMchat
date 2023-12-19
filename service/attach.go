package service

import (
	"fmt"
	"ginchat/utils"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var FileName string
func GetPicName(pic string) string {
	return pic
}
//文件上床
func Upload(c *gin.Context) { // 上传单个文件
	//write := c.ResponseWriter
	w 	   := c.Writer
	req	   := c.Request
	src,head,err := req.FormFile("file") //file 是前端传过来的key
	if err != nil {
		utils.RepFail(w,err.Error())
	}
	suffix := ".png" //默认后缀
	ofiname := head.Filename  //原始文件名
	t := strings.Split(ofiname,".") //分割文件名
	if len(t) > 1 { 
		suffix = "."+t[len(t)-1] //获取后缀
	}
	fileName := fmt.Sprintf("%d%04d%s", time.Now().Unix(),rand.Int31(),suffix) //生成文件名,格式是时间戳+随机数+后缀，例如：1600000000+1234+.png
	dst,err := os.Create("./asset/upload/"+fileName) //创建文件
	if err!= nil {
		utils.RepFail(w,err.Error()) 
	}
	_ , err  = io.Copy(dst, src)
	if err!= nil {
		utils.RepFail(w,err.Error())
	}
	FileName = GetPicName(fileName)
	url := "./asset/upload/"+fileName
	utils.RepOK(w,url,"发送图片成功")
}

