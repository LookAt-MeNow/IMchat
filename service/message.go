package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

//防止跨域站点的伪请求
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { 
		return true
	},
}

//SendMsg
func SendMsg(c *gin.Context) {
	ws,err := upgrader.Upgrade(c.Writer,c.Request,nil) //升级get请求为webSocket协议
	if err!= nil {
		fmt.Println("Failed to upgrade connection:", err)
		fmt.Println(err)
		return	
	}
	defer func (ws *websocket.Conn)  {
		err := ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)
	MsgHandle(ws, c) //处理消息
}

func MsgHandle(ws *websocket.Conn, c *gin.Context){ 
	msg,err := utils.Subscribe(c,utils.PublishKey) //订阅redis
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("sendMsg success")
	tm := time.Now().Format("2006-01-02 15:04:05") //获取当前时间
	m := fmt.Sprintf("[ws][%s]: %s",tm,msg)        //拼接消息
	err = ws.WriteMessage(1,[]byte(m))             //发送消息
	if err != nil {
		fmt.Println(err)
		return
	}
}

func  Chat(c *gin.Context)  {
	models.Chat(c.Writer,c.Request)
}