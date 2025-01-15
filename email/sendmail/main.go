package main

import (
	"crypto/tls"
	"fmt"
	"html/template"
	"net/smtp"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	// 邮箱的 SMTP 服务器地址和端口
	//smtpServer := "smtp-mail.outlook.com"
	//smtpPort := 587

	smtpServer := "10.1.234.44"
	smtpPort := 25
	// 发件人的邮箱地址和密码
	//from := "499968985@qq.com"
	username := "mailtest"
	password := "Qwert@12345"

	// 收件人的邮箱地址
	from := "499968985@qq.com"
	to := "xiongyang@noahgroup.com"

	// 邮件主题
	subject := "Test Email"

	// 认证信息
	auth := smtp.PlainAuth("", username, password, smtpServer)

	// TLS 配置
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer,
	}

	// 连接到 SMTP 服务器
	conn, err := tls.Dial("tcp", smtpServer+":"+fmt.Sprint(smtpPort), tlsConfig)
	if err != nil {
		fmt.Println("Error connecting to SMTP server:", err)
		return
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, smtpServer)
	if err != nil {
		fmt.Println("Error creating SMTP client:", err)
		return
	}
	defer client.Close()

	// 身份验证
	if err = client.Auth(auth); err != nil {
		fmt.Println("Error authenticating:", err)
		return
	}

	// 设置发件人和收件人
	if err = client.Mail(from); err != nil {
		fmt.Println("Error setting sender:", err)
		return
	}
	if err = client.Rcpt(to); err != nil {
		fmt.Println("Error setting recipient:", err)
		return
	}

	// 发送邮件数据
	w, err := client.Data()
	if err != nil {
		fmt.Println("Error getting data writer:", err)
		return
	}
	defer w.Close()
	dates := new(struct {
		Value string
	})
	////赋值
	dates.Value = "123456"
	body := FormatEmailBody("./view/verfityemail.html", dates)
	_, err = w.Write([]byte("From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
		"\r\n" +
		body + "\r\n"))
	if err != nil {
		fmt.Println("Error writing email data:", err)
		return
	}

	fmt.Println("Email sent successfully")
}

func FormatEmailBody(path string, data interface{}) string {
	//创建一个字符串缓冲
	builder := &strings.Builder{}
	//定义模板函数
	funcs := map[string]interface{}{
		"dateformat": func(t *time.Time) string {
			if t == nil {
				return ""
			}
			return t.Format("2006-01-02 15:04:05")
		},
	}
	//使用template.New创建一个模板对象，并使用ParseFiles方法将模板文件解析为模板对象
	tpl := template.Must(template.New("tpl").Funcs(funcs).ParseFiles(path))
	//使用模板对象执行并写入到字符串中
	err := tpl.ExecuteTemplate(builder, filepath.Base(path), data)
	if err != nil {
		return err.Error()
	}
	//返回字符串
	return builder.String()
}
