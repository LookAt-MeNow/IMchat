package models

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"sync"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	FormId 		uint	//发送者
	TargetId 	uint	//接受者
	Type 		int	//消息类型	
	Media		int		//消息种类
	Content		string	//消息内容
	CreateTime 	uint64 	//创建时间
	Pic			string	//图片
	Url			string	//地址
	Desc		string	
	Amount		int
}

func (table *Message) TableName() string {
	return "message"
}

type Node struct {
	Conn 		*websocket.Conn
	DataQueu 	chan []byte
	GroupSets 	set.Interface
}

//映射关系
var clientMap map[uint] *Node = make(map[uint] *Node,0)
//读写锁
var Locker sync.RWMutex

func Chat(wr http.ResponseWriter,req *http.Request) {
	//获取参数
	query 	 := req.URL.Query()
	//token  := query.Get("token")
	id  	 := query.Get("userId")
	userid,_ := strconv.ParseInt(id,10,32)
	//msgtype  := query.Get("type")
	//context  := query.Get("context")
	//检验token
	isright := true
	conn,err := (&websocket.Upgrader{
		CheckOrigin: func (r *http.Request) bool  {
			return isright
		},
	}).Upgrade(wr,req,nil)
	if err!= nil {
		fmt.Println(err)
		return
	}
	//获取连接
	node := &Node{
		Conn: 		conn,
		DataQueu: 	make(chan []byte,50),
		GroupSets:  set.New(set.ThreadSafe),
	}
	//用户关系
	//userid 跟 node绑定
	Locker.Lock()
	clientMap[uint(userid)] = node
	Locker.Unlock()
	//发送逻辑
	go sendProc(node)
	//接收逻辑
	go recvProc(node)
	sendMsg(uint(userid),[]byte("欢迎进入！！"))
}

func sendProc(node *Node) {
	for data := range node.DataQueu {
		err := node.Conn.WriteMessage(websocket.TextMessage,data)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func recvProc(node *Node) {
	for {
		_, data ,err := node.Conn.ReadMessage() 
			if err != nil {
				fmt.Println(err)
				return
			}
		broadMsg(data)
		fmt.Println("[ws]...",string(data))
	}
}

var upsendchan chan []byte = make(chan []byte,1024)
func broadMsg(data []byte) {
	upsendchan <- data
}

func init(){
	go upSendProc()
	go upRecvProc()
	fmt.Println("Inti goroutine...")
}


//完成upd数据发送协程
func upSendProc() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 255),
		Port: 3030,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if cerr := conn.Close(); cerr != nil {
			fmt.Println(cerr)
		}
	}()
	for data := range upsendchan {
		_, err := conn.Write(data)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

//完成udp数据发送协程
func upRecvProc() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3030,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if cerr := conn.Close(); cerr != nil {
			fmt.Println(cerr)
		}
	}()
	for {
		var buf [1024]byte
		n,err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		//解析数据
		fmt.Println("udp data...",string(buf[0:n]))
		dispatch(buf[0:n])
	}
}

//后端调度逻辑处理
func dispatch(data []byte) {
	msg := Message{}
	//msg. = uint64(time.Now().Unix())
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch msg.Type {
	case  1: //私信
		fmt.Println("dispatch  data :", string(data))
		sendMsg( msg.TargetId, data)
	//case "2": //群发
		//sendGroupMsg(msg.TargetId, data) //发送的群ID ，消息内容
		// case "4": // 心跳
		// 	node.Heartbeat()
		//case "4":
		//
	}
}

func sendMsg(useid uint,msg []byte) {
	Locker.Lock()
	node,ok := clientMap[useid]
	Locker.Unlock()
	if ok {
		node.DataQueu <- msg
	}
}