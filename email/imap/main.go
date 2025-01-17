package main

import (
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"log"
)

func main() {
	// 设置 IMAP 服务器地址和端口
	imapServer := "partner.outlook.cn:993"

	// 设置邮箱地址和密码
	username := "guidang@noahfound.partner.onmschina.cn"
	password := "Nuv54390"

	// 连接到 IMAP 服务器
	c, err := client.DialTLS(imapServer, nil)
	if err != nil {
		log.Fatal("conn ", err)
	}
	defer c.Logout()

	// 登录邮箱
	if err := c.Login(username, password); err != nil {
		log.Fatal("login: ", err)
	}

	// 选择邮箱邮箱
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Total:", mbox.Messages)

	// 读取邮件
	seqset := new(imap.SeqSet)
	seqset.AddRange(1, mbox.Messages)
	messages := make(chan *imap.Message, 10)
	done := make(chan error, 1)
	//section := &imap.BodySectionName{}
	go func() {
		//done <- c.Fetch(seqset, []imap.FetchItem{section.FetchItem()}, messages)
		done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
	}()

	//
	//Get the whole message body

	for msg := range messages {
		log.Println("Message:", msg.Envelope.Subject)
		//r := msg.GetBody(section)
		//if r == nil {
		//	log.Fatal("Server didn't returned message body")
		//}
		//m, err := mail.ReadMessage(r)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//header := m.Header
		////log.Println("Date:", header.Get("Date"))
		////log.Println("From:", header.Get("From"))
		////log.Println("To:", header.Get("To"))
		//log.Println("Subject:", header.Get("Subject"))
		//var buf bytes.Buffer
		//buf.ReadFrom(m.Body)
		//log.Println("Body:", buf.String())
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}
}
