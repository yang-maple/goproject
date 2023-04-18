package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
)

// 定消息类型的结构体
type Msg struct {
	User       string // 用户名
	Msgsession string // 消息体
}

// 定义客户端类型的结构体
type Client struct {
	Addr    string      // ip地址
	Session chan []byte // 接收消息管道
}

// 接收管道信息
var msgchan = make(chan Msg)

// 客户端存储信息
// map[addr]{addr,msg}
var custerip = make(map[string]Client)

// 接收写入管道函数
func getmsg() {
	for {
		// 循环进行信息的获取,如果msgchan 没有数据将会一直堵塞
		// 将获取信息写入msgs
		msgs := <-msgchan
		for _, custers := range custerip {
			msgBytes, _ := json.Marshal(msgs)
			// 将信息写入管道
			custers.Session <- msgBytes
		}
	}

}

func read(conn net.Conn) {
	for {
		// 定义接受数据的类型为字节
		data := make([]byte, 1024)
		//读取客户端数据
		n, err := conn.Read(data)
		if err == nil {
			// 获捕获客户端信息并进行处理
			readmsg := strings.TrimSpace(string(data[:n]))
			fmt.Println("接收到数据：", string(readmsg))
			msgchan <- Msg{User: conn.RemoteAddr().String(), Msgsession: readmsg}
		}
	}
}
func send(conn net.Conn, cclient Client) {

	for {
		// 管道数据进行赋值
		// 循环向每个客户端发送数据
		msgBytes := <-cclient.Session
		fmt.Printf("已向%s发送:%s\n", cclient.Addr, string(msgBytes))
		_, _ = conn.Write(msgBytes)
	}
}

// 处理函数
func Handle(conn net.Conn) {
	log.Printf("%s 已连接成功", conn.RemoteAddr().String())
	// 给变量赋值
	custer := Client{conn.RemoteAddr().String(), make(chan []byte)}
	custerip[conn.RemoteAddr().String()] = custer
	// 对客户端发来的信息进行解析读取
	go read(conn)
	// 向客户端进行广播的发送操作
	go send(conn, custer)
	//

}

func main() {
	// 服务ip和端口
	addr := "0.0.0.0:8081"
	//  监听消息
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%s 已经开启", addr)
	// 延迟关闭监听
	defer listen.Close()
	go getmsg()

	for {
		// 接受连接
		conn, err := listen.Accept()
		if err == nil {
			// 对所有的监听进行处理
			go Handle(conn)
		}
	}

}
