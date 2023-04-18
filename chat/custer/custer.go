package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
)

type Msg struct {
	User       string // 用户名
	Msgsession string // 消息体
}

// 接收消息的变量
var Session Msg

func read(conn net.Conn) {
	for {
		// 定义接收数据的切片
		data := make([]byte, 1024)
		// 接收数据
		n, err := conn.Read(data)
		if err == nil {
			// 对接收数据进行处理
			// 去除空格输入
			readmsg := strings.TrimSpace(string(data[:n]))
			// 使用Msg 类型变量反序列化接收数据
			json.Unmarshal([]byte(readmsg), &Session)
			// 输入数据
			fmt.Printf("\n%s:%s\n", Session.User, Session.Msgsession)
		}

	}
}

func write(conn net.Conn) {
	for {
		var message string
		fmt.Print("you:")
		fmt.Scan(&message)
		// 对写入的数据 进行转化
		jsonmsg, err := json.Marshal(message)
		if err == nil {
			//fmt.Println("已发送：", string(jsonmsg))
			// 写入数据
			conn.Write(jsonmsg)
		}

	}

}

func connect(conn net.Conn) {
	// 延迟关闭连接
	defer conn.Close()
	// 并发处理客户端 读操作

	go read(conn)

	// 处理客户端的 写操作
	write(conn)

}

func main() {
	// 访问ip 端口
	addr := "127.0.0.1:8081"
	// 进行 tcp 连接
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%s 已连接上\n", addr)
	// 对 连接 进行处理
	connect(conn)

}
