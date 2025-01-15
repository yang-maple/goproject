package main

import (
	"fmt"
	"github.com/bytbox/go-pop3"
	"log"
)

func main() {
	// 设置邮箱服务器地址和端口
	server := "mail.noahgroup.com:995"

	// 设置邮箱地址和密码
	username := "xiy19425_admin"
	password := "20200725.zyj"

	// 连接到 POP3 服务器
	client, err := pop3.DialTLS(server)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Quit()

	// 登录邮箱
	if err := client.User(username); err != nil {
		log.Fatal(err)
	}
	if err := client.Pass(password); err != nil {
		log.Fatal(err)
	}

	// 获取邮箱中的邮件数量
	count, size, err := client.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Mailbox has %d messages with a total size of %d bytes\n", count, size)

	// 读取邮件

}
