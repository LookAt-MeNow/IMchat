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

//文件上床
func Upload(c *gin.Context) {
	//write := c.ResponseWriter
	w 	   := c.Writer
	req	   := c.Request
	src,head,err := req.FormFile("file")
	if err != nil {
		utils.RepFail(w,err.Error())
	}
	suffix := ".png"
	ofiname := head.Filename
	t := strings.Split(ofiname,".")
	if len(t) > 1 {
		suffix = "."+t[len(t)-1]
	}
	fileName := fmt.Sprintf("%d%04d%s", time.Now().Unix(),rand.Int31(),suffix)
	dst,err := os.Create("./asset/upload/"+fileName)
	if err!= nil {
		utils.RepFail(w,err.Error())
	}
	_ , err  = io.Copy(dst, src)
	if err!= nil {
		utils.RepFail(w,err.Error())
	}
	url := "./asset/upload/"+fileName
	utils.RepOK(w,url,"发送图片成功")
}